package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	ID int
}

func (e *Event) consume() {
	time.Sleep(1 * time.Millisecond)
}

const (
	numThreads = 100
	bufferSize = 20
	numLoops   = 10
)

func producer(theBuffer chan<- *Event, numLoops int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numLoops; i++ {
		e := &Event{ID: i}
		theBuffer <- e
	}
}

func consumer(theBuffer <-chan *Event, wg *sync.WaitGroup) {
	defer wg.Done()

	for e := range theBuffer {
		e.consume()
	}
}

func main() {
	aBuffer := make(chan *Event, bufferSize)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	numProducers := numThreads / 2
	numConsumers := numThreads - numProducers

	producerWg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go producer(aBuffer, numLoops, &producerWg)
	}

	consumerWg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go consumer(aBuffer, &consumerWg)
	}

	producerWg.Wait()

	close(aBuffer)

	consumerWg.Wait()

	fmt.Println("All producers and consumers finished successfully.")
}
