package generator

import (
	"encoding/json"
	"engine/generation"
	"engine/metadata"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"sync"
)

func ConsumeMetadata(subsciber <-chan metadata.Metadata, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		go consumer(subsciber, wg)
	}
}

func consumer(subsciber <-chan metadata.Metadata, wg *sync.WaitGroup) {
	for md := range subsciber {
		generateImage(md)
		saveMetadata(md)
		wg.Done()
	}
}

func saveMetadata(md metadata.Metadata) {
	path := fmt.Sprintf("build/metadata/%d.json", md.Id)

	out, err := os.Create(path)
	if err != nil {
		panic("Metadata could not be saved")
	}
	defer out.Close()

	res, err := json.Marshal(md)
	if err != nil {
		panic("Could not convert to json")
	}

	ioutil.WriteFile(path, res, os.ModePerm)
}

func generateImage(md metadata.Metadata) {
	dest := image.NewRGBA(image.Rect(0, 0, generation.Size, generation.Size))

	fmt.Println("Generating image ", md.Id)

	for _, attribute := range md.Attributes {
		im := open(attribute.Path)
		addLayer(dest, im)
	}

	path := fmt.Sprintf("build/images/%d.png", md.Id)

	out, err := os.Create(path)
	if err != nil {
		panic("image was not created")
	}
	defer out.Close()

	png.Encode(out, dest)
}

func addLayer(dest draw.Image, src image.Image) {
	draw.Draw(dest, dest.Bounds(), src, image.Point{}, draw.Over)
}

func open(path string) image.Image {
	srcF, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Could not open path %s", path))
	}
	defer srcF.Close()

	src, err := png.Decode(srcF)
	if err != nil {
		panic(fmt.Sprintf("Could not decode image at path %s", path))
	}

	return src
}
