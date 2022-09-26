package main

import (
	"engine/configuration"
	"engine/metadata"
	"engine/save"
	"engine/timer"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	defer timer.Timer(fmt.Sprintf("Generating took"))()

	config := configuration.Parse()

	totalCount := config.Generations[len(config.Generations)-1].GrowEditionSizeTo

	mdc := make(chan *metadata.Metadata)

	wg := &sync.WaitGroup{}
	wg.Add(totalCount)

	go save.MetadataHandler(config, mdc, func() {
		wg.Done()
	})

	metadata.StartGeneration(config, mdc)

	wg.Wait()
}
