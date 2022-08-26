package metadata

import (
	"engine/generation"
	"fmt"
	"math/rand"
)

type Attribute struct {
	Path      string `json:"-"`
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type Metadata struct {
	Id         int         `json:"-"`
	Name       string      `json:"name"`
	Image      string      `json:"image"`
	Attributes []Attribute `json:"attributes"`
}

var dna = make(map[string]int) // maps DNA to the id of the generated metadata

func StartGeneration(count int, path string, layers []string, listener chan<- Metadata) {
	// TODO: move initLayers into generation steps nested for loop
	initLayers(layers, path)

	list := rand.Perm(count)

	for _, id := range list {
		metadata := generateMetadata(id)

		for !isUnique(metadata) {
			fmt.Println("DNA is not unique")
			metadata = generateMetadata(id)
		}

		// res, err := json.Marshal(metadata)
		// if err != nil {
		// 	panic("Could not convert to json")
		// }

		// fmt.Println(string(res))

		listener <- metadata
	}
}

func generateMetadata(id int) Metadata {
	attributes := []Attribute{}

	for _, layer := range layers {
		attr := layer.RandomItem()
		attributes = append(attributes, attr)
	}

	name := fmt.Sprintf("%s #%d", generation.Name, id)
	image := generation.Image

	return Metadata{Id: id, Name: name, Image: image, Attributes: attributes}
}

func isUnique(metadata Metadata) bool {
	hash := ""
	for _, attribute := range metadata.Attributes {
		hash += fmt.Sprintf("%s:%s-", attribute.TraitType, attribute.Value)
	}

	v, found := dna[hash]

	if found { // dna is not unique
		return false
	}

	dna[hash] = v // save dna in the map to check later

	return true
}
