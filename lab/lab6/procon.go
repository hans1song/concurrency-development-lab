// Package main demonstrates the Producer-Consumer concurrency pattern.
// It uses Go channels for communication and Sync.WaitGroups for synchronization.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Event represents a unit of work to be processed.
type Event struct {
	ID int
}

// consume simulates processing the event.
func (e *Event) consume() {
	time.Sleep(1 * time.Millisecond)
}

const (
	numThreads = 100
	bufferSize = 20
	numLoops   = 10
)

// producer generates events and sends them to the buffer channel.
func producer(theBuffer chan<- *Event, numLoops int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numLoops; i++ {
		e := &Event{ID: i}
		theBuffer <- e
	}
}

// consumer reads events from the buffer channel and processes them.
// It continues until the channel is closed and drained.
func consumer(theBuffer <-chan *Event, wg *sync.WaitGroup) {
	defer wg.Done()

	for e := range theBuffer {
		e.consume()
	}
}

func main() {
	// Create a buffered channel to act as the shared queue.
	aBuffer := make(chan *Event, bufferSize)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	numProducers := numThreads / 2
	numConsumers := numThreads - numProducers

	// Start producer goroutines
	producerWg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go producer(aBuffer, numLoops, &producerWg)
	}

	// Start consumer goroutines
	consumerWg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go consumer(aBuffer, &consumerWg)
	}

	// Wait for all producers to finish sending events.
	producerWg.Wait()

	// Close the channel to signal consumers that no more events will come.
	close(aBuffer)

	// Wait for all consumers to finish processing the remaining events.
	consumerWg.Wait()

	fmt.Println("All producers and consumers finished successfully.")
}
