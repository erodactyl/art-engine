package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type LayerItem struct {
	Name        string
	Path        string
	Coefficient int
}

type Layer struct {
	Name            string
	DisplayName     string
	IgnoreDNA       bool
	Optionality     float32
	Items           []*LayerItem
	CoefficientsSum int
}

type Generation struct {
	GrowEditionSizeTo int
	Size              int
	Layers            []*Layer
}

type Configuration struct {
	Name        string
	ImageBase   string
	Size        int
	Generations []*Generation
}

func Parse() *Configuration {
	content, err := ioutil.ReadFile("./config.json")

	if err != nil {
		panic(fmt.Sprintf("Can't open the config.json file due to %s", err))
	}

	var config Configuration

	err = json.Unmarshal(content, &config)

	if err != nil {
		panic(fmt.Sprintf("Can't open the config.json file due to %s", err))
	}

	prev := 0
	for _, generationStep := range config.Generations {
		generationStep.Size = generationStep.GrowEditionSizeTo - prev
		prev = generationStep.GrowEditionSizeTo

		for _, layer := range generationStep.Layers {
			if layer.DisplayName == "" {
				layer.DisplayName = layer.Name
			}

			setupLayerItems(layer)
		}
	}

	return &config
}

func (layer Layer) RandomItem() *LayerItem {
	// Return nil for optionality
	test := rand.Float32()
	if test < layer.Optionality {
		return nil
	}

	index := rand.Intn(layer.CoefficientsSum) + 1

	for _, item := range layer.Items {
		index -= item.Coefficient
		if index <= 0 {
			return item
		}
	}

	panic("Mistake in layer item randomization")
}

func setupLayerItems(layer *Layer) {
	dir, err := os.Open(fmt.Sprintf("%s/%s", "layers", layer.Name))
	if err != nil {
		panic(fmt.Sprintf("Could not setup layer %s", layer.Name))
	}

	files, err := dir.Readdir(0)
	if err != nil {
		panic(fmt.Sprintf("Could not read files for layer %s", layer.Name))
	}

	layerItems := []*LayerItem{}
	coefficientsSum := 0

	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".png") {
			path := fmt.Sprintf("layers/%s/%s", layer.Name, fileName)
			name, coefficient := getNameAndCoeffient(fileName)

			item := &LayerItem{Path: path, Name: name, Coefficient: coefficient}
			layerItems = append(layerItems, item)

			coefficientsSum += coefficient
		}
	}

	layer.Items = layerItems
	layer.CoefficientsSum = coefficientsSum

	fmt.Println(len(layer.Items))
}

func getNameAndCoeffient(fileName string) (string, int) {
	parts := strings.Split(fileName, "#")

	if len(parts) == 1 {
		name := strings.TrimSuffix(parts[0], ".png")

		return name, 1
	}

	name := parts[0]

	coefficientStr := strings.TrimSuffix(parts[1], ".png")
	coefficient, err := strconv.Atoi(coefficientStr)

	if err != nil {
		panic(fmt.Sprintf("Can't parse coefficient for file %s", fileName))
	}

	return name, coefficient
}
