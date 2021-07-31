package datapipeline

import (
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

// pipe 1
func (dp *dataPipeline) parseCSVData(done <-chan struct{}, filepath string) (<-chan string, <-chan error) {
	emailChan := make(chan string)
	errChan := make(chan error, 1)

	go func() {
		defer close(emailChan)

		file, err := os.Open(filepath)
		if file != nil {
			defer file.Close()
		}
		if err != nil {
			errChan <- err
			return
		}

		if !fileTypeIsCSV(file) {
			errChan <- errors.New("file provided to parseCSVData must be of type text/csv")
			return
		}

		if _, err := file.Seek(0, 0); err != nil {
			errChan <- err
			return
		}

		reader := csv.NewReader(file)
		emailIndex := getEmailIndex(*reader)
		if emailIndex < 0 {
			errChan <- errors.New("email field not found")
			return
		}

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				continue
			}

			email := record[emailIndex]

			select {
			case emailChan <- email:
			case <-done:
				return
			}
		}
	}()

	return emailChan, errChan
}

func fileTypeIsCSV(file io.Reader) bool {
	mt, _ := mimetype.DetectReader(file)

	return mt.Is("text/csv")
}

func getEmailIndex(cr csv.Reader) (idx int) {
	idx = -1
	fields, err := cr.Read()
	if err != nil {
		return
	}

	for i, f := range fields {
		if f != "email" {
			continue
		}

		idx = i
	}

	return
}
