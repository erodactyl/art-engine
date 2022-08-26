package main

import (
	"engine/generation"
	"engine/generator"
	"engine/metadata"
	"engine/timer"
	"fmt"
	"path/filepath"
	"sync"
)

func main() {
	defer timer.Timer(fmt.Sprintf("Generating %d images took", generation.Count))()

	basePath, err := filepath.Abs(generation.LayersPath)
	if err != nil {
		panic(fmt.Sprintf("Count not locate basePath at %s", generation.LayersPath))
	}

	c := make(chan metadata.Metadata)

	wg := sync.WaitGroup{}
	wg.Add(generation.Count)

	go generator.ConsumeMetadata(c, &wg)

	metadata.StartGeneration(generation.Count, basePath, generation.Layers, c)

	wg.Wait()
}
