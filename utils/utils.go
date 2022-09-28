package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func ParallelForEach[T any](elements []T, operation func(el T), routines int) {
	busy := make(chan bool, routines)

	for _, _el := range elements {
		busy <- true
		go (func(el T) {
			operation(el)
			<-busy
		})(_el)
	}
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
