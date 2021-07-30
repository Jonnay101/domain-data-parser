package datapipeline

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

// pipe 1
func (dp *dataPipeline) parseCSVData(filepath string) <-chan string {
	emailChan := make(chan string)

	go func() {
		file, err := os.Open(filepath)
		if file != nil {
			defer file.Close()
		}
		if err != nil {
			log.Fatal(err)
		}

		if !fileTypeIsCSV(file) {
			log.Fatal("file provided to parseCSVData must be of type text/csv")
		}

		if _, err := file.Seek(0, 0); err != nil {
			log.Fatal(err)
		}

		reader := csv.NewReader(file)
		emailIndex := getEmailIndex(*reader)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				continue
			}

			email := record[emailIndex]

			emailChan <- email
		}

		close(emailChan)
	}()

	return emailChan
}

func fileTypeIsCSV(file io.Reader) bool {
	mt, _ := mimetype.DetectReader(file)

	return mt.Is("text/csv")
}

func getEmailIndex(cr csv.Reader) (idx int) {
	fields, err := cr.Read()
	if err != nil {
		log.Fatal(err)
	}

	idx = -1
	for i, f := range fields {
		if f != "email" {
			continue
		}

		idx = i
	}

	if idx < 0 {
		log.Fatal("email field title not found")
	}

	return
}
