package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type inputFile struct {
	filepath  string
	seperator string
	pretty    bool
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func check(e error) {
	if e != nil {
		exitGracefully((e))
	}
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

func processCsvFile(fileData inputFile, writerChannel chan<- map[string]string) {
	file, err := os.Open(fileData.filepath)
	check(err)
	defer file.Close()

	var headers, line []string
	reader := csv.NewReader(file)

	if fileData.separator == "semicolon" {
		reader.Comma = ';'
	}

	headers, err = reader.Read()

	check(err)

	for {
		line, err = reader.Read()

		if err == io.EOF {
			close(writerChannel)
			break
		} else if err != nil {
			exitGracefully(err)
		}

		record, err := processLine(headers, line)

		if err != nil {
			fmt.Printf("Line: %sError :%s\n", line, err)
			continue
		}

		writerChannel <- record
	}

}

func main() {
	// fileData, err := getFileData()
}
