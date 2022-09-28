package main

import (
	"engine/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		panic("please provide an IPFS hash")
	}

	hash := os.Args[1]

	if len(hash) != 46 {
		panic("please provide a valid IPFS hash")
	}

	dir, err := os.Open("build/metadata")
	if err != nil {
		panic("Could not load metadata directory")
	}

	_files, err := dir.Readdir(0)
	if err != nil {
		panic("Could not read metadata files")
	}

	files := []os.FileInfo{}
	for _, file := range _files {
		if strings.HasSuffix(file.Name(), ".json") {
			files = append(files, file)
		}
	}

	utils.ParallelForEach(files, func(file os.FileInfo) {
		path := fmt.Sprintf("build/metadata/%s", file.Name())
		updateFile(path, hash)

		fmt.Println(fmt.Sprintf("File %s was updated", path))

	}, 10)

	fmt.Println(fmt.Sprintf("%d files updated", len(files)))
}

func updateFile(path string, hash string) {
	path, err := filepath.Abs(path)
	if err != nil {
		panic(fmt.Sprintf("could not update file %s", path))
	}

	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("could not update file %s", path))
	}

	updated := strings.Replace(string(content), "REPLACE", hash, 1)

	if err := os.WriteFile(path, []byte(updated), 0666); err != nil {
		panic(fmt.Sprintf("could not update file %s", path))
	}
}
