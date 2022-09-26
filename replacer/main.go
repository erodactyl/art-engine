package main

func main() {
	panic("Not updated")
	// if len(os.Args) == 1 {
	// 	panic("please provide an IPFS hash")
	// }
	//
	// hash := os.Args[1]
	//
	// if len(hash) != 46 {
	// 	panic("please provide a valid IPFS hash")
	// }
	//
	// wg := sync.WaitGroup{}
	// wg.Add(generation.Count)
	//
	// c := make(chan int, 10)
	//
	// for i := 0; i < generation.Count; i++ {
	// 	c <- i
	// 	go func(i int) {
	// 		updateFile(i+1, hash)
	// 		fmt.Println("done")
	// 		<-c
	// 		wg.Done()
	// 	}(i)
	// }
	//
	// wg.Wait()
}

// func updateFile(id int, hash string) {
// 	path, err := filepath.Abs(fmt.Sprintf("build/metadata/%d.json", id))
// 	if err != nil {
// 		panic(fmt.Sprintf("could not update file %d", id))
// 	}
//
// 	content, err := os.ReadFile(path)
// 	if err != nil {
// 		panic(fmt.Sprintf("could not update file %d", id))
// 	}
//
// 	updated := strings.Replace(string(content), "REPLACE", hash, 1)
//
// 	if err := os.WriteFile(path, []byte(updated), 0666); err != nil {
// 		panic(fmt.Sprintf("could not update file %d", id))
// 	}
// }
