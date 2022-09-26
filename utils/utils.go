package utils

import (
	"encoding/json"
	"fmt"
)

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
		index := list[currentIndex]
		currentIndex++
		return index
	}
}

func randomize(count int) []int {
	list := make([]int, count)

	for i := 0; i < count; i++ {
		list[i] = i + 1
	}

	return list
}
