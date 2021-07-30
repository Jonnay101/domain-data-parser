package datapipeline

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

func Test_fileTypeIsCSV(t *testing.T) {
	falseyCSV, err := os.Open("test_files/test.txt")
	if err != nil {
		t.Error(err)
	}
	defer falseyCSV.Close()

	truthyCSV, err := os.Open("test_files/titles_missing.csv")
	if err != nil {
		t.Error(err)
	}
	defer truthyCSV.Close()

	type args struct {
		file io.Reader
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not a CSV file", args{falseyCSV}, false},
		{"is a CSV file", args{truthyCSV}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileTypeIsCSV(tt.args.file); got != tt.want {
				t.Errorf("fileTypeIsCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEmailIndex(t *testing.T) {
	hasFields, err := os.Open("test_files/big_test_data.csv")
	if err != nil {
		t.Error(err)
	}
	defer hasFields.Close()

	hasFieldsReader := csv.NewReader(hasFields)

	noFields, err := os.Open("test_files/titles_missing.csv")
	if err != nil {
		t.Error(err)
	}
	defer noFields.Close()

	noFieldsReader := csv.NewReader(noFields)

	type args struct {
		rdr *csv.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"not a CSV file", args{hasFieldsReader}, 2},
		{"is a CSV file", args{noFieldsReader}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEmailIndex(*tt.args.rdr); got != tt.want {
				t.Errorf("getEmailIndex() = %d, want %d", got, tt.want)
			}
		})
	}
}
