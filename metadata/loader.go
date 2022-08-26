package metadata

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var layers = []Layer{}

type Layer struct {
	TraitType string
	Items     []Attribute
}

func (l Layer) RandomItem() Attribute {
	index := rand.Intn(len(l.Items))
	attribute := l.Items[index]
	return attribute
}

func initLayers(layerNames []string, basePath string) {
	// reset layers from previous generation step
	layers = []Layer{}

	for _, layerName := range layerNames {
		layer := setupLayer(layerName, basePath)
		layers = append(layers, layer)
	}
}

func setupLayer(name string, basePath string) Layer {
	dir, err := os.Open(fmt.Sprintf("%s/%s", basePath, name))
	if err != nil {
		panic(fmt.Sprintf("Could not setup layer %s at path %s/", name, basePath))
	}

	files, err := dir.Readdir(0)
	if err != nil {
		panic(fmt.Sprintf("Could not read files for layer %s at path %s/", name, basePath))
	}

	layerItems := []Attribute{}

	for i := range files {
		fileName := files[i].Name()
		if strings.HasSuffix(fileName, ".png") {
			path := fmt.Sprintf("layers/%s/%s", name, fileName)
			attribute := Attribute{Path: path, TraitType: name, Value: fileName}
			layerItems = append(layerItems, attribute)
		}
	}

	return Layer{TraitType: name, Items: layerItems}
}
