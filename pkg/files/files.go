package files

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetListFileWithPrefix(path, prefix string) ([]string, error) {
	filesWithPrefix := make([]string, 0)
	listFile, err := getListFileFromDir("./resources")
	if err != nil {
		return filesWithPrefix, err
	}

	for _, f := range listFile {
		if strings.HasPrefix(f, prefix) {
			filesWithPrefix = append(filesWithPrefix, f)
		}
	}
	return filesWithPrefix, nil
}

func getListFileFromDir(path string) ([]string, error) {
	listFile := make([]string, 0)
	listDir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return listFile, err
	}

	for _, dir := range listDir {
		listFile = append(listFile, dir.Name())
	}
	return listFile, nil
}

func GetFileContents(filePath string) ([]string, error) {
	contents := make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return contents, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	return contents, nil
}
