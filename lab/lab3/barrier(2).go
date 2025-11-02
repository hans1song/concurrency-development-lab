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
// Issues:
// The barrier is not implemented!
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var (
	totalRoutines = 10
	mutex         sync.Mutex
	count         int
	sem           *semaphore.Weighted
	ctx           = context.TODO()
)

// barrier implemented with semaphore
func barrier() {
	mutex.Lock()
	count++
	if count == totalRoutines {

		sem.Release(int64(totalRoutines))
	} else {

		mutex.Unlock()
		sem.Acquire(ctx, 1)
		return
	}
	mutex.Unlock()
}

func doStuff(goNum int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	barrier()

	fmt.Println("Part B", goNum)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	sem = semaphore.NewWeighted(int64(totalRoutines))
	sem.Acquire(ctx, int64(totalRoutines))
	for i := 0; i < totalRoutines; i++ {

		go doStuff(i, &wg)
	}

	wg.Wait()
	fmt.Println("All done!")
}
