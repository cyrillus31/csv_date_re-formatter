package main

import (
	"fmt"

	"github.com/cyrillus31/csv_date_re-formatter/utils"
)

const INPUT_FOLDER = "input_files/"
const OUTPUT_FOLDER = "output_files/"

func main() {
	files := utils.FindAllInputFiles(INPUT_FOLDER)
	// fmt.Printf("%v\n", files)
	for _, file := range files {
		// utils.PrintFileContent(file)
		result, _ := utils.GetFileContent(file)
	}

}
