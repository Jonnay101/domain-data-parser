package datapipeline

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// pipe 1
func (dp *dataPipeline) parseCSVData(filepath string) <-chan string {
	emailChan := make(chan string)

	// TODO: run a csv file type check

	go func() {
		file, err := os.Open(filepath)
		if file != nil {
			defer file.Close()
		}
		if err != nil {
			log.Fatal(err)
		}

		reader := csv.NewReader(file)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				continue
			}

			email := record[2] // TODO: detect email idx from field names

			emailChan <- email
		}

		close(emailChan)
	}()

	return emailChan
}
