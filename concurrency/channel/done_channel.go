package channel

import (
	"fmt"
	"time"
)

func DoneChannel() {
	//use close channel
	done := make(chan interface{})
	//array
	in := []interface{}{1, 2, 3, 4, 5}
	//stream processing
	simpleStream := func(done <-chan interface{}, in ...interface{}) <-chan interface{} {
		results := make(chan interface{})
		go func() {
			defer close(results)
			defer fmt.Println("goroutine closed")
			for _, v := range in {
				select {
				case results <- v:
					//sleep 1 second
					time.Sleep(time.Second)
				case <-done:
					return
				}
			}
		}()
		return results
	}
	//use(print) return result
	consumer := func(done chan interface{}, results <-chan interface{}) {
		for v := range results {
			//close doneChannel
			if v.(int) == 2 {
				close(done)
			}
			fmt.Println(v)
		}
	}
	results := simpleStream(done, in...)
	consumer(done, results)
}
