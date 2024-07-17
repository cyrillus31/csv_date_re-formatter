package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func FindAllInputFiles(inputFolder string) []string {
	files := []string{}
	filepath.WalkDir(inputFolder, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func GetFileContent(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, errors.New("ERROR: An eror occured when tried to opan a file")
	}
	result := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), "\t"))
	}
	// fmt.Println(result)
	return result, nil
}

func PrintFileContent(filePath string) error {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return errors.New("ERROR: An eror occured when tried to opan a file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(strings.Split(scanner.Text(), "\t")[:])
	}
	return nil
}
