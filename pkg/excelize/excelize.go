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

func PopulateSheetsName(fileName, sheetName string, keyColumn, searchColumn []string) error {
	f, err := excelize.OpenFile("./resources/" + fileName)
	if err != nil {
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	data, err := generateKeyFromAllSheets(f, sheetName, searchColumn)
	if err != nil {
		return err
	}
	for i := 2; i < 83007; i++ {
		indexKey := make([]string, 0)
		for _, key := range keyColumn {
			val, _ := f.GetCellValue(sheetName, key+strconv.Itoa(i))
			indexKey = append(indexKey, val)
		}

		ik := strings.Join(indexKey, "##")
		fn := make([]string, 0)
		uk := make([]string, 0)
		for _, dik := range data[ik] {
			d := strings.Split(dik, "##")
			fn = append(fn, d[0])
			uk = append(uk, d[1])
		}
		f.SetCellValue(sheetName, "H"+strconv.Itoa(i), strings.Join(fn, ", "))
		f.SetCellValue(sheetName, "I"+strconv.Itoa(i), strings.Join(uk, ", "))
	}
	f.Save()
	return nil
}

func generateKeyFromAllSheets(f *excelize.File, sheetName string, searchColumn []string) (map[string][]string, error) {
	sheets := f.GetSheetList()
	data := make(map[string][]string, 0)
	for _, sheet := range sheets {
		if sheet != sheetName {
			last := false
			for i := 2; !last; i++ {
				indexKey := make([]string, 0)
				for _, key := range searchColumn {
					val, _ := f.GetCellValue(sheet, key+strconv.Itoa(i))
					if val == "" {
						last = true
						break
					}
					indexKey = append(indexKey, val)
				}

				ik := strings.Join(indexKey, "##")
				_, ok := data[ik]
				if !ok {
					data[ik] = make([]string, 0)
				}
				uniqueKey, _ := f.GetCellValue(sheet, "F"+strconv.Itoa(i))
				data[ik] = append(data[ik], sheet+"##"+uniqueKey)
			}
		}
	}
	return data, nil
}
