//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by:
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type ReusableBarrier struct {
	n       int32
	count   atomic.Int32
	mutex   sync.Mutex
	enterCh chan struct{}
	leaveCh chan struct{}
}

func NewReusableBarrier(n int) *ReusableBarrier {
	return &ReusableBarrier{
		n:       int32(n),
		enterCh: make(chan struct{}),
		leaveCh: make(chan struct{}),
	}
}

func (b *ReusableBarrier) Wait() {
	b.mutex.Lock()
	enterChannel := b.enterCh
	leaveChannel := b.leaveCh
	b.mutex.Unlock()

	if b.count.Add(1) == b.n {
		close(enterChannel)
	}

	<-enterChannel

	if b.count.Add(-1) == 0 {
		b.mutex.Lock()
		b.enterCh = make(chan struct{})
		b.leaveCh = make(chan struct{})
		b.mutex.Unlock()

		close(leaveChannel)
	}

	<-leaveChannel
}

func doStuff(goNum int, wg *sync.WaitGroup, barrier *ReusableBarrier) {
	defer wg.Done()

	fmt.Printf("Part A - Goroutine %d working\n", goNum)
	time.Sleep(time.Second)
	fmt.Printf("Part A - Goroutine %d reached the barrier\n", goNum)

	barrier.Wait()

	fmt.Printf("\nPart B - Goroutine %d working\n", goNum)
	time.Sleep(time.Second)
	fmt.Printf("Part B - Goroutine %d reached the second barrier\n", goNum)

	barrier.Wait()

	fmt.Printf("\nPart C - Goroutine %d finished\n", goNum)
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	barrier := NewReusableBarrier(totalRoutines)

	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &wg, barrier)
	}
	wg.Wait()
	fmt.Println("All done!")
}
