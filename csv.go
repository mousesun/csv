package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CsvReader struct {
	Reader *csv.Reader
}

func NewCsv(filepath string) (*CsvReader, error) {
	file, err := os.Open(filepath)
	if checkError(err) {
		return nil, err
	}

	reader := csv.NewReader(file)

	var m_reader CsvReader
	m_reader.Reader = reader
	return &m_reader, nil
}

func (reader *CsvReader) Next() ([]string, error) {
	line, err := reader.Reader.Read()
	if checkError(err) {
		return nil, err
	}

	return line, nil
}

func (reader *CsvReader) NextRows(rows int) (records [][]string, err error) {
	count := 0
	for {
		if count >= rows {
			return records, nil
		}

		record, err := reader.Reader.Read()
		if err == io.EOF {
			return records, nil
		}

		if err != nil {
			return nil, err
		}

		records = append(records, record)
		count = count + 1
	}
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println("Error:", err)
		return true
	}

	return false
}
