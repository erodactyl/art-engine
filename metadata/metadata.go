package metadata

import (
	"engine/configuration"
	"engine/utils"
	"fmt"
)

type Attribute struct {
	Path      string `json:"-"`
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type Metadata struct {
	Id         int          `json:"-"`
	Name       string       `json:"name"`
	Image      string       `json:"image"`
	Attributes []*Attribute `json:"attributes"`
}

type context struct {
	getIndex  func() int
	config    *configuration.Configuration
	dna       map[string]int
	metadatas []*Metadata
	mc        chan<- *Metadata
}

func StartGeneration(config *configuration.Configuration, mc chan<- *Metadata) {
	totalCount := config.Generations[len(config.Generations)-1].GrowEditionSizeTo
	getIndex := utils.IndexFactory(totalCount)

	dna := make(map[string]int) // maps DNA to the id of the generated metadata

	ctx := &context{getIndex, config, dna, []*Metadata{}, mc}

	for _, generation := range config.Generations {
		generationStep(generation, ctx)
	}
}

func generationStep(generation *configuration.Generation, ctx *context) {
	for i := 0; i < generation.Size; i++ {
		index := ctx.getIndex()

		metadata := generateMetadata(index, generation, ctx)

		for !isUnique(metadata, ctx) {
			fmt.Println("DNA is not unique")
			metadata = generateMetadata(index, generation, ctx)
		}

		ctx.mc <- metadata
	}
}

func randomAttributeFromLayer(layer *configuration.Layer) *Attribute {
	item := layer.RandomItem()

	if item == nil {
		return nil
	}

	attribute := &Attribute{Path: item.Path, TraitType: layer.Name, Value: item.Name}
	return attribute
}

func generateMetadata(id int, generation *configuration.Generation, ctx *context) *Metadata {
	attributes := []*Attribute{}

	for _, layer := range generation.Layers {
		attr := randomAttributeFromLayer(layer)

		if attr != nil {
			attributes = append(attributes, attr)
		}
	}

	name := fmt.Sprintf("%s #%d", ctx.config.Name, id)
	image := fmt.Sprintf("ipfs://%s/%d.png", ctx.config.ImageBase, id)

	metadata := &Metadata{Id: id, Name: name, Image: image, Attributes: attributes}

	ctx.metadatas = append(ctx.metadatas, metadata)

	return metadata
}

func isUnique(metadata *Metadata, ctx *context) bool {
	hash := ""
	for _, attribute := range metadata.Attributes {
		hash += fmt.Sprintf("%s:%s-", attribute.TraitType, attribute.Value)
	}

	v, found := ctx.dna[hash]

	if found { // dna is not unique
		return false
	}

	ctx.dna[hash] = v // save dna in the map to check later

	return true
}
