package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	os.RemoveAll("./dataset")
	os.Mkdir("./dataset", 0755)
	os.MkdirAll("./dataset/train/images", 0755)
	os.MkdirAll("./dataset/train/labels", 0755)
	os.MkdirAll("./dataset/val/images", 0755)
	os.MkdirAll("./dataset/val/labels", 0755)

	files, err := os.ReadDir("./test")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
