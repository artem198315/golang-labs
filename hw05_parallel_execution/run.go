package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type errorCount struct {
	mu sync.Mutex
	i  int
}

func errorHandle(wg *sync.WaitGroup, chErr chan error, stopCh chan struct{}, errCount *errorCount, maxErrors int) {
	defer wg.Done()

	for err := range chErr {
		fmt.Println(err)
		errCount.mu.Lock()
		errCount.i++
		errCount.mu.Unlock()
		if errCount.i >= maxErrors {
			fmt.Printf("%d %d\n", errCount.i, maxErrors)
			close(stopCh)
			return
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskCh := make(chan Task, len(tasks))
	chErr := make(chan error, m)
	stopCh := make(chan struct{}, 1)
	errCount := errorCount{}

	for _, i := range tasks {
		taskCh <- i
	}
	close(taskCh)

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(worker int) {
			defer wg.Done()
			for f := range taskCh {
				select {
				case <-stopCh:
					fmt.Println("Stop worker ", worker)
					return
				default:
				}

				select {
				case <-stopCh:
					fmt.Println("Stop worker ", worker)
					return
				default:
					fmt.Println("Worker ", worker)
					if err := f(); err != nil {
						chErr <- err
					}
				}
			}
		}(i)
	}

	wgErr := sync.WaitGroup{}
	wgErr.Add(1)
	go errorHandle(&wgErr, chErr, stopCh, &errCount, m)

	wg.Wait()
	close(chErr)

	fmt.Println("Tasks completed")
	wgErr.Wait()

	if errCount.i >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
