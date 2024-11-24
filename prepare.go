package main

import (
	"fmt"
	"io"
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

	dirs := []string{"train", "val"}
	for _, dir := range dirs {

		files, err := os.ReadDir("./" + dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			label := file.Name()
			fmt.Println(file.Name())
			files2, _ := os.ReadDir("./" + dir + "/" + label)
			for _, file2 := range files2 {
				fmt.Println(file2.Name())
				copy("./"+dir+"/"+label+"/"+file2.Name(), "./dataset/"+dir+"/images/ "+label+"_"+file2.Name())
				os.WriteFile("./dataset/"+dir+"/labels/"+label+"_"+file2.Name()+".txt", []byte(label), 0644)

			}

		}
	}

}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
