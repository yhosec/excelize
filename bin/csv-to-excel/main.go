package main

import (
	"excelize/pkg/excelize"
	"excelize/pkg/files"
	"flag"
)

func main() {
	var prefix = flag.String("prefix", "t", "type your prefix of csv file")
	var separator = flag.String("separator", ",", "type your separator of csv file")
	flag.Parse()
	dir := "./resources"
	listFile, err := files.GetListFileWithPrefix(dir, *prefix)
	if err != nil {
		return
	}

	csv := make(map[string][]string, 0)
	for _, f := range listFile {
		content, err := files.GetFileContents("./resources" + "/" + f)
		if err != nil {
			return
		}

		csv[f] = content
	}

	excelize.GenerateExcelFromCsv(csv, listFile, *separator, *prefix)
}
