package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidParameters   = errors.New("invalid parameters provided")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func beforeValidator(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return ErrInvalidParameters
	}

	if n <= 0 {
		return ErrInvalidParameters
	}

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func inputTasks(ch chan Task, tasks []Task) {
	for i := 0; i < len(tasks); i++ {
		ch <- tasks[i]
	}
}

func Run(tasks []Task, n int, m int) error {
	if err := beforeValidator(tasks, n, m); err != nil {
		if errors.Is(err, ErrErrorsLimitExceeded) {
			return err
		}
		return nil
	}

	tasksCh := make(chan Task, len(tasks))
	inputTasks(tasksCh, tasks)

	var (
		errCnt int64
		wg     sync.WaitGroup
	)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				t, ok := <-tasksCh
				if !ok {
					return
				}

				if err := t(); err != nil {
					atomic.AddInt64(&errCnt, 1)
				}

				if atomic.LoadInt64(&errCnt) >= int64(m) {
					return
				}
			}
		}()
	}
	close(tasksCh)

	wg.Wait()

	if errCnt >= int64(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
