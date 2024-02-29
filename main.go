package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := flag.String("path", "./dir", "path to scan folder")
	upper := flag.Bool("upper", false, "to upper case")
	json := flag.Bool("json", false, "json output")
	flag.Parse()
	if *path == "./" || *path == "../" || *path == ".." {
		log.Println("Wrong path!")
		return
	}
	if *json {
		fmt.Println(toJson(changeCase(scanDir(*path), *upper)))
	} else {
		for k, v := range changeCase(scanDir(*path), *upper) {
			fmt.Println(k, ">", v)
		}
	}
}

func toJson(array map[string]string) string {
	jsonStr, err := json.Marshal(array)
	if err != nil {
		log.Println(err)
	}
	return string(jsonStr)
}

func scanDir(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files = append(files, path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return files
}

func changeCase(fileNames []string, upper bool) map[string]string {
	files := make(map[string]string)
	for _, fileName := range fileNames {
		var newFileName string
		if upper {
			newFileName = strings.ToUpper(fileName)
		} else {
			newFileName = strings.ToLower(fileName)
		}
		if fileName != "main.go" && fileName != "go.mod" && fileName != "." && fileName != "./" && fileName != "../" && fileName != ".." {
			os.Rename(fileName, newFileName)
			files[fileName] = newFileName
		}
	}
	return files
}
