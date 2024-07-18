package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
  "time"
)


func isInSlice(target int, slice *[]int) bool {
  for _, value := range *slice {
    if value == target {
      return true
    }
  }
  return false
}

type Table struct {
  Filename        string
	Date            string
	Time            string
	DateTime        string
  DateTimeLayout  string
	RawTable        [][]string
  DateRowNumber   int
  RowNumberToKeep []int
}

func InitializeTable(filename string, tableContent [][]string) Table {
	return Table{
    Filename: filename,
		Date:     "",
		Time:     "",
		DateTime: "",
		RawTable: tableContent,
    DateRowNumber: 0,
	}
}

func (t *Table) GetDateRowNumber() int {
	fmt.Print("Укажите букву столбца с датой, который нужно отформатировать: ")
	reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
  result := int(input[0] - 'a')
  t.DateRowNumber = result
	return result
}

func (t *Table) ConvertData() [][]string {
  var convertedData = [][]string{}
  for i := 0; i < len(t.RawTable); i++ {
    row := t.RawTable[i]
    var newRow = []string{}
    for index, value := range row {
      // fmt.Println(newRow)
      if index == t.DateRowNumber {
        date, time := DateConverter(value)
        newRow = append(newRow, date, time)
        continue
      } else if isInSlice(index, &t.RowNumberToKeep) {
        newRow = append(newRow, value)
      }
    }
    convertedData = append(convertedData, newRow)
  }
  return convertedData
}

func (t *Table) GetRowNumbers() []int {
	var input string
	fmt.Print("Введите латинские буквы, соответствующие нужным столбцам:\n")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.Trim(input, " \n\t")
	return func() []int {
		var result = []int{}
		arr := strings.Split(input, " ")
		// fmt.Println(arr)
		for _, element := range arr {
			result = append(result, int(element[0]-'a'))
		}
    t.RowNumberToKeep = result
		return result
	}()

}

func DateConverter(inputDateTime string) (string, string) {
  inputLayout := "2006-01-02 15:04:05" 
  datetime, _ := time.Parse(inputLayout, inputDateTime)
  time := datetime.Format("15:04")
  date := datetime.Format("2/1/2006")
  return date, time
}


func FindAllInputFiles(inputFolder string) []string {
	files := []string{}
	filepath.WalkDir(inputFolder, func(path string, d fs.DirEntry, err error) error {
    fileName := d.Name()
		if !d.IsDir() && fileName[0] != '.' {
			files = append(files, fileName)
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
  return result[1:len(result)], nil
}

func printFileContent(filePath string) error {
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
