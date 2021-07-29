package emaildomainstats

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func parseCSVData(filepath string) <-chan string {
	emailChan := make(chan string)

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

			email := record[2]

			emailChan <- email
		}

		close(emailChan)
	}()

	return emailChan
}
