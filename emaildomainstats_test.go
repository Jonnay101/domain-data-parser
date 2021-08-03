/*
This package is required to provide functionality to process the supplied customer_data.csv file
and return a sorted (by email domain) data structure of your choice containing the email domains along with the number of
customers for each domain. Any errors should be logged (or handled) or returned to the consumer of this package.
Performance matters, the sample file may only contains 1K lines but it could be 1m lines or run on a small machine.
*/

package emaildomainstats

import (
	"bytes"
	"encoding/csv"
	"io"
	"reflect"
	"strings"
	"testing"
)

func createCSVFile(rows ...string) (io.Reader, error) {
	records := make([][]string, len(rows))
	buff := bytes.NewBuffer([]byte{})
	w := csv.NewWriter(buff)

	for _, row := range rows {
		rowSplit := strings.Split(row, ",")
		records = append(records, rowSplit)
	}

	w.WriteAll(records)

	return buff, nil
}

func BenchmarkGetDomainNameStats1000Lines(b *testing.B) {
	benchmarkGetDomainNameStats("customer_data.csv", b)
}

func BenchmarkGetDomainNameStats8Lines(b *testing.B) {
	benchmarkGetDomainNameStats("test_customer_data1.csv", b)
}

func benchmarkGetDomainNameStats(filepath string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetDomainNameStats(filepath)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func TestGetDomainNameStats(t *testing.T) {
	wantStore := Store{
		stats: []DomainStat{
			{"about.com", 1},
			{"bigcartel.com", 2},
			{"goo.gl", 1},
			{"google.com", 2},
			{"google.com.jp", 1},
		},
		indexReference: map[string]int{
			"about.com":     0,
			"bigcartel.com": 1,
			"goo.gl":        2,
			"google.com":    3,
			"google.com.jp": 4,
		},
	}
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    Store
		wantErr bool
	}{
		{"parse customer_data.csv", args{"test_customer_data1.csv"}, wantStore, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDomainNameStats(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_parseEmailsFromFile(b *testing.B) {
	reader, err := createCSVFile(
		"first_name,last_name,email,gender,ip_address",
		"Rowland,Aldins,raldins0@google.com,Female,46.148.249.225",
		"Zaria,Brugh,zbrugh1@goo.gl,Bigender,235.106.185.201",
		"Gamaliel,Burgoine,gburgoine2@bigcartel.com,Non-binary,76.214.44.117",
		"Godfree,Chattelaine,gchattelaine3@about.com,Polygender,73.139.226.59",
		"Karlis,Le Clercq,kleclercq4@google.com,Female,58.114.31.41",
		"Ximenes,Sumbler,xsumbler5@google.com.jp,Male,238.161.80.177",
		"James,Struder,jsjs2@bigcartel.com,Male,83.192.43.109",
	)
	if err != nil {
		b.Error(err)
		return
	}

	for n := 0; n < b.N; n++ {
		_, err := parseEmailsFromFile(reader)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func Test_parseEmailsFromFile(t *testing.T) {
	emptyRes := make([]string, 0)
	validRes := []string{"email@zero.com", "email@apple.io", "email@xerox.co"}

	emailColumnMissing, err := createCSVFile("name,gender", "too,short")
	if err != nil {
		t.Error(err)
		return
	}
	wrongNumFields, err := createCSVFile("name,gender,email,age", "two,fields")
	if err != nil {
		t.Error(err)
		return
	}
	invalidEmail, err := createCSVFile("name,gender,email,age", "col,non-binary,not-email,22")
	if err != nil {
		t.Error(err)
		return
	}
	validEmails, err := createCSVFile(
		"name,gender,email,age",
		"Col,non-binary,email@zero.com,22",
		"Sue,female,email@apple.io,42",
		"Dave,male,email@xerox.co,19",
	)
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		f io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"emailColumnMissing", args{emailColumnMissing}, emptyRes, true},
		{"wrong number of fields", args{wrongNumFields}, emptyRes, true},
		{"invalid email", args{invalidEmail}, emptyRes, false},
		{"valid email list", args{validEmails}, validRes, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseEmailsFromFile(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEmailsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEmailsFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_parseDomainsFromEmailX4(b *testing.B) {
	emails := []string{"email@zero.com", "email@apple.io", "email2@apple.io", "email@xerox.co"}

	benchmark_parseDomainsFromEmail(emails, b)
}

func benchmark_parseDomainsFromEmail(emails []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseDomainsFromEmail(emails)
	}
}

func Test_parseDomainsFromEmail(t *testing.T) {
	// testList := []string{"email@zero.com", "email@apple.io", "email2@apple.io", "email@xerox.co"}
	// wantList := []string{"zero.com", "apple.io", "apple.io", "xerox.co"}

	testInvalidEmailList := []string{"email£apple.io", "email@xerox.co"}
	wantInvalidEmailList := []string{"zero.com", "xerox.co"}
	type args struct {
		emails []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// {"ordered list", args{testList}, wantList},
		{"invalid email in list", args{testInvalidEmailList}, wantInvalidEmailList},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDomainsFromEmail(tt.args.emails); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDomainsFromEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_storeSortableDomainStats(b *testing.B) {
	domains := []string{"xerox.co", "apple.io", "apple.io", "zero.com", "zero.com", "zero.com"}

	for i := 1; i < 15; i++ {
		i++
		doms := domains
		domains = append(domains, doms...)
	}

	b.Logf("%d domains sent to storeSortableDomainStats", len(domains))

	store := Store{}

	for n := 0; n < b.N; n++ {
		if err := store.storeSortableDomainStats(domains); err != nil {
			b.Error(err)
			return
		}
	}
}

func Test_storeSortableDomainStats(t *testing.T) {
	testList := []string{"xerox.co", "apple.io", "apple.io", "zero.com", "zero.com", "zero.com"}

	wantStore := Store{
		stats: []DomainStat{
			{"apple.io", 2},
			{"xerox.co", 1},
			{"zero.com", 3},
		},
		indexReference: map[string]int{
			"apple.io": 0,
			"xerox.co": 1,
			"zero.com": 2,
		},
	}

	type args struct {
		domains []string
	}
	tests := []struct {
		name    string
		args    args
		want    Store
		wantErr bool
	}{
		{"should be cool", args{testList}, wantStore, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := Store{}
			err := store.storeSortableDomainStats(tt.args.domains)
			if (err != nil) != tt.wantErr {
				t.Errorf("storeEmails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(store, tt.want) {
				t.Errorf("storeEmails() = %v, want %v", store.GetAllStats(), tt.want)
			}
		})
	}
}

func TestStore_GetDomainCountByName(t *testing.T) {
	tesStore := Store{
		stats: []DomainStat{
			{"apple.io", 2},
			{"xerox.co", 1},
			{"zero.com", 3},
		},
		indexReference: map[string]int{
			"apple.io": 0,
			"xerox.co": 1,
			"zero.com": 2,
		},
	}
	type args struct {
		domainName string
	}
	tests := []struct {
		name  string
		store Store
		args  args
		want  int32
	}{
		{"apple.io is there", tesStore, args{"apple.io"}, 2},
		{"hotmail.com isn't in there", tesStore, args{"hotmail.com"}, 0},
		{"xerox.co is there", tesStore, args{"xerox.co"}, 1},
		{"zero.com is there", tesStore, args{"zero.com"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.store.GetDomainCountByName(tt.args.domainName); got != tt.want {
				t.Errorf("Store.GetDomainCountByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
