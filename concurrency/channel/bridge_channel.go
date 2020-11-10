package channel

import (
	"fmt"
	"sync"
)

func BridgeChannel() {
	done := make(chan interface{})
	defer close(done)
	wg := sync.WaitGroup{}
	//generator
	generator := func(done <-chan interface{}, in ...interface{}) <-chan interface{} {
		inStream := make(chan interface{})
		go func() {
			defer close(inStream)
			for _, v := range in {
				select {
				case inStream <- v:
				case <-done:
					return
				}
			}
		}()
		return inStream
	}

	chan1 := generator(done, 1, 2, 3, 4, 5)
	chan2 := generator(done, 6, 7, 8, 9, 10)
	chanStream := make(chan (<-chan interface{}))

	wg.Add(2)
	//Add chanStream
	go func() {
		chanStream <- chan1
		chanStream <- chan2
	}()

	bridge := func(done <-chan interface{}, chanStream <-chan (<-chan interface{})) <-chan interface{} {
		results := make(chan interface{})
		go func() {
			for {
				var stream <-chan interface{}
				select {
				case <-done:
					return
				case stream = <-chanStream:
				}
				//Loop to 1 channel
				go func() {
					defer wg.Done()
					for v := range stream {
						select {
						case <-done:
							return
						default:
							results <- v
						}
					}
				}()
			}
		}()
		//Await all channel success
		go func() {
			wg.Wait()
			close(results)
		}()
		return results
	}

	results := bridge(done, chanStream)
	for v := range results {
		fmt.Println(v)
	}
}
