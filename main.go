package main

import (
	// "encoding/csv"
	// "fmt"
	// "log"
	// "bufio"
	// "os"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	// "strings"
)

const INPUT_FOLDER = "input_files/"
const OUTPUT_FOLDER = "output_files/"

func findAllInputFiles(inputFolder string) []string {
	files := []string{}
	filepath.WalkDir(inputFolder, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func printFileContent(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Println("ERROR: An eror occured when tried to opan a file", err)
	}
}

func main() {
	files := findAllInputFiles(INPUT_FOLDER)
	fmt.Printf("%v\n", files)
	// file, err := os.Open(filepath.Join(INPUT_FOLDER, ))

}
