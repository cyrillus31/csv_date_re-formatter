package main

import (
	// "fmt"

	"fmt"
	"path"
	"time"

	"github.com/cyrillus31/csv_date_re-formatter/utils"
)

const INPUT_FOLDER = "input_files/"
const OUTPUT_FOLDER = "output_files/"

func main() {
	time.Sleep(3 * time.Second)
	files := utils.FindAllInputFiles(INPUT_FOLDER)
	fmt.Println(files)
	for _, file := range files {
		fileContent, _ := utils.GetFileContent(path.Join(INPUT_FOLDER, file))
		table := utils.InitializeTable(file, fileContent)
    table.GetRowNumbers()
    table.GetDateRowNumber()
    newData := table.ConvertData()
    fmt.Println(newData)
	}

}
