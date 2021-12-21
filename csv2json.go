package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type inputFile struct {
	filepath  string
	seperator string
	pretty    bool
}

func getFileData() (inputFile, error) {
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("A filepath argument is required")
	}

	seperator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse()

	fileLocation := flag.Arg(0)

	if !(*seperator == "comma" || *seperator == "semicolon") {
		return inputFile{}, errors.New("Only comma or semicolon seperators are allowed")
	}

	return inputFile{fileLocation, *seperator, *pretty}, nil

}

func checkIfValidFile(filename string) (bool, error) {
	if fileExtension := filepath.Ext(filename); fileExtension != ".csv" {
		return false, fmt.Errorf("File %s is not CSV", filename)
	}

	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %s does not exist", filename)
	}

	return true, nil
}

func main() {
	// fileData, err := getFileData()
}
