package main

import (
	"excelize/pkg/excelize"
	"flag"
	"strings"
)

func main() {
	var fileName = flag.String("file", "", "type target file name from resources folder")
	var mainSheet = flag.String("sheet", "", "type sheet name from main data")
	var keyColumn = flag.String("key", "", "type column will used as key")
	var searchColumn = flag.String("search", "", "type column will used for search by key")
	flag.Parse()

	excelize.PopulateSheetsName(
		*fileName,
		*mainSheet,
		strings.Split(*keyColumn, ","),
		strings.Split(*searchColumn, ","),
	)
}
