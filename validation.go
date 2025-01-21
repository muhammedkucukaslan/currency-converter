package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func validateFlags(sourceFile string, base string, target string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Amount can not be 0 or less")
	}

	//* maybe i could do this in main.go
	file, _ := os.Open(sourceFile)
	mymap, err := CSVToMap(file)
	if err != nil {
		return fmt.Errorf("Formating map:%v", err)
	}

	if mymap[strings.ToUpper(base)] == "" {
		return fmt.Errorf("Invalid base symbol.")
	}

	if mymap[strings.ToUpper(target)] == "" {
		return fmt.Errorf("Invalid target symbol.")
	}

	return nil
}

// * https://gist.github.com/muhammedkucukaslan/73f01d6628c6b842d5ece84fc3ab084f
func CSVToMap(reader io.Reader) (map[string]string, error) {
	csvReader := csv.NewReader(reader)

	header, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	if len(header) < 2 {
		return nil, fmt.Errorf("CSV file must have at least two columns")
	}

	targetMap := make(map[string]string)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read record: %w", err)
		}

		targetMap[record[0]] = record[1]
	}

	return targetMap, nil
}
