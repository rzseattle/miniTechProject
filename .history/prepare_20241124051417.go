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
		label := file.Name()
		fmt.Println(file.Name())
		files2, _ := os.ReadDir("./test/" + label)
		for _, file2 := range files2 {
			fmt.Println(file2.Name())
		}

	}

}
