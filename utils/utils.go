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

var rowMap = map[rune]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
}

func GetRowNumbers() []int {
	var input string
	fmt.Print("Введите латинские буквы, соответствующие нужным столбцам:\n")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	return func() []int {
		var result = []int{}
		arr := strings.Split(input, " ")
		// fmt.Println(arr)
		for _, element := range arr {
			result = append(result, int(element[0]-'a'))
		}
		return result
	}()

}

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
