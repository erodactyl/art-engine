package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
)

func ParallelForEach[T any](elements []T, operation func(el T), routines int) {
	// Wait group to wait until all operations are ran before moving on
	wg := sync.WaitGroup{}
	wg.Add(len(elements))

	busy := make(chan bool, routines)

	for _, _el := range elements {
		// Add operation to the queue, `routines` max
		busy <- true

		go (func(el T) {
			operation(el)

			// Mark operation done
			wg.Done()

			// Operation is done, free the queue
			<-busy
		})(_el)
	}

	wg.Wait()
}

func DeepPrint(object interface{}) {
	objectJSON, err := json.MarshalIndent(&object, "", "  ")

	if err == nil {
		fmt.Println(string(objectJSON))
	}
}

func IndexFactory(count int) func() int {
	list := randomize(count)
	currentIndex := 0

	return func() int {
		id := list[currentIndex]
		currentIndex++

		return id
	}
}

func randomize(count int) []int {
	list := rand.Perm(count)

	for i := 0; i < count; i++ {
		list[i] = list[i] + 1
	}

	return list
}
