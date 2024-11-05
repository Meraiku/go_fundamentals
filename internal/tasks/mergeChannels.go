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
				for {
					select {
					case v, ok := <-ch:
						if !ok {
							return
						}
						inCh <- v
					}
				}
			}()

		}
		wg.Wait()
		close(inCh)

	}()

	go func() {
		defer close(outCh)

		for {
			select {
			case v, ok := <-inCh:
				if !ok {
					return
				}
				outCh <- v
			}
		}

	}()

	return outCh
}
