package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

var fileToReadFrom = "files/p2data.txt"
var fileToWriteTo = "plot/p2data.csv"
var fileSeparator = " "

func main() {
	txt := readTxt(fileToReadFrom)
	writeCsv(txt, fileToWriteTo, fileSeparator)
}

func readTxt(filename string) []string {
	var result []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func writeCsv(textToWrite []string, filename string, fileSeparator string) {
	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	for i := 0; i < len(textToWrite); i++ {
		row := strings.Split(textToWrite[i], fileSeparator)[1:5] //delete first element of each entry

		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
