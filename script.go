package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
	txt := readTxt("files/results.txt")
	writeCsv(txt, "results/records.csv", " ")
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
	records := textToWrite

	file, err := os.Create(filename)

	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, record := range records {
		row := strings.Split(record, fileSeparator)
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
