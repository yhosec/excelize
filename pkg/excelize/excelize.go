package excelize

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func GenerateExcelFromCsv(csv map[string][]string, sheets []string, separator, fileName string) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for h := 0; h < len(sheets); h++ {
		// for fileName, bulkData := range csv {
		indexSheet, _ := f.NewSheet(sheets[h])
		f.SetActiveSheet(indexSheet)
		for i, rowData := range csv[sheets[h]] {
			for j, data := range strings.Split(rowData, separator) {
				// fmt.Println(j, toCharStr(j)+strconv.Itoa(i+1), data)
				f.SetCellValue(sheets[h], toCharStr(j)+strconv.Itoa(i+1), data)
			}
		}
	}

	err := f.SaveAs("./resources/output/" + fileName + ".xlsx")
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func toCharStr(i int) string {
	return arr[i]
}
