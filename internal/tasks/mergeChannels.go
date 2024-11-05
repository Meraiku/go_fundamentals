package tasks

import (
	"runtime"
	"sync"
)

func MergeChannels[T any](channels ...<-chan T) <-chan T {

	inCh := make(chan T, runtime.NumCPU())
	outCh := make(chan T)

	wg := &sync.WaitGroup{}

	go func() {
		for _, ch := range channels {
			wg.Add(1)

			go func() {
				defer wg.Done()
				for v := range ch {
					inCh <- v
				}
			}()

		}
		wg.Wait()
		close(inCh)

	}()

	go func() {
		defer close(outCh)

		for v := range inCh {
			outCh <- v
		}

	}()

	return outCh
}
