package splitsies

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// Splitter ...
type Splitter struct {
	WithHeader           bool
	FileNamePrefix       string
	FileNameColumnIndex  int
	ColumnValueMaxLength int
	InputCSVPath         string
	OutputDirPath        string
}

// Split splits file in smaller chunks
func (s Splitter) Split() ([]string, error) {
	file, err := os.Open(s.InputCSVPath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", s.InputCSVPath)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.ReuseRecord = true

	firstRecord := true
	var results []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if firstRecord {
			firstRecord = false
			totalCols := len(record)
			if s.FileNameColumnIndex < 0 || s.FileNameColumnIndex >= totalCols {
				return nil, errors.New("file-name-column-index is out of range")
			}

			if s.WithHeader {
				continue
			}
		}

		fnColValue := record[s.FileNameColumnIndex]
		if s.ColumnValueMaxLength > 0 && len(fnColValue) > s.ColumnValueMaxLength {
			fnColValue = fnColValue[0:s.ColumnValueMaxLength]
		}

		writeToFile := fmt.Sprintf("%s/%s%s.csv", s.OutputDirPath, s.FileNamePrefix, fnColValue)
		f, err := os.OpenFile(writeToFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return results, err
		}

		writer := csv.NewWriter(f)
		if err := writer.Write(record); err != nil {
			return results, err
		}
		writer.Flush()
		f.Close()
	}

	return nil, nil
}
