package main

import (
	// "fmt"

	"encoding/csv"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/cyrillus31/csv_date_re-formatter/utils"
)

const INPUT_FOLDER = "input_files/"
const OUTPUT_FOLDER = "output_files/"


func changeExt2CSV(fileName string) string {
  ext := path.Ext(fileName)
  lenWithoutExtenstion := len(fileName) - len(ext)
  return fileName[:lenWithoutExtenstion] + ".csv"
}

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
    // fmt.Println(fileContent[:20])
    // fmt.Println(newData[:20])
    file := changeExt2CSV(file)
    f, _ := os.Create(path.Join(OUTPUT_FOLDER, file)) 
    w := csv.NewWriter(f)
    
    for _, row := range newData {
      w.Write(row)
    }
	}

}
