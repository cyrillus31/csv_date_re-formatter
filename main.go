package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/cyrillus31/csv_date_re-formatter/utils"
)

const INPUT_FOLDER = "input_files"
const OUTPUT_FOLDER = "output_files"
var empty string


func changeExt2CSV(fileName string) string {
  ext := path.Ext(fileName)
  lenWithoutExtenstion := len(fileName) - len(ext)
  return fileName[:lenWithoutExtenstion] + ".csv"
}

func at_startup() {
  os.MkdirAll(INPUT_FOLDER, os.ModePerm)
  os.MkdirAll(OUTPUT_FOLDER, os.ModePerm)
  fmt.Printf("В папке с запущенной программой появилась директория %v. Сложите в нее все файлы одного формата, требующие конвертации, после чего нажмите Enter.\n", INPUT_FOLDER)
  fmt.Scanln(&empty)
}

func at_finish() {
  fmt.Printf("Перефоматированные файлы были помещены в директорию %v. Работа программы будет завершена.\nНажмите Enter чтобы выйти.", OUTPUT_FOLDER)
  fmt.Scanln(&empty)    

  cmd := "open"
  split := "/"
  if runtime.GOOS == "windows" {
      cmd = "explorer"
      split = "\\"
  }
  cwd, _ := os.Getwd()
  // fullPath := path.Join(cwd, OUTPUT_FOLDER)
  fullPath := cwd + split + OUTPUT_FOLDER
  fmt.Println(fullPath)
  // exec.Command(cmd, "\\", fullPath).Start()
  exec.Command(cmd, fullPath).Start()
}

func main() {
  at_startup()
  defer at_finish()

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
