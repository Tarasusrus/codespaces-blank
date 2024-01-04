package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files{
		filePath := filepath.Join(path, file.Name())
		if file.IsDir(){
			scanDirectory(filePath)
	} else {
		fmt.Println(filePath)
		}
	}
}	

func calmDown() {
	recover()
}

func one() {
	defer fmt.Println("defer in one")
	two()
}

func two() {
	defer fmt.Println("defer in two")
	three()
}

func three() {
	defer calmDown()
	panic("unimplemented")
}
func main() {
	defer fmt.Println("defer in main")
	scanDirectory("my_directory")
	one()
}