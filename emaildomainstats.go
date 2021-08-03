/*
This package is required to provide functionality to process the supplied customer_data.csv file
and return a sorted (by email domain) data structure of your choice containing the email domains along with the number of
customers for each domain. Any errors should be logged (or handled) or returned to the consumer of this package.
Performance matters, the sample file may only contains 1K lines but it could be 1m lines or run on a small machine.
*/

package emaildomainstats

import (
	"encoding/csv"
	"errors"
	"io"
	"net/mail"
	"os"
	"sort"
	"strings"
)

func GetDomainNameStats(filepath string) (Store, error) {
	store := Store{}

	file, err := os.Open(filepath)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		return store, err
	}

	emails, err := parseEmailsFromFile(file)
	if err != nil {
		return store, err
	}

	domains := parseDomainsFromEmail(emails)

	return store, store.storeSortableDomainStats(domains)
}

func parseEmailsFromFile(f io.Reader) ([]string, error) {
	emails := make([]string, 0)
	emailIdx := -1
	r := csv.NewReader(f)

	rowCount := 0

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return emails, err
		}

		if rowCount == 0 {
			emailIdx = findColumnIndex("email", row)
			if emailIdx < 0 {
				return emails, errors.New("could not find `email` column header")
			}
		}

		rowCount++

		if len(row) < (emailIdx + 1) {
			continue
		}

		if _, err := mail.ParseAddress(row[emailIdx]); err != nil {
			continue
		}

		emails = append(emails, row[emailIdx])
	}

	return emails, nil
}

func findColumnIndex(wantColumn string, row []string) int {
	idx := -1

	for i, cn := range row {
		if wantColumn != cn {
			continue
		}

		idx = i
	}

	return idx
}

func parseDomainsFromEmail(emails []string) []string {
	domainList := make([]string, 0)

	for _, email := range emails {
		emailSlice := strings.Split(email, "@")
		if len(emailSlice) < 2 {
			continue
		}

		domainList = append(domainList, emailSlice[1])

	}

	return domainList
}

type DomainStat struct {
	Name  string
	Count int32
}

type Store struct {
	stats          []DomainStat
	indexReference map[string]int
}

func (s *Store) storeSortableDomainStats(domains []string) error {
	s.indexReference = make(map[string]int)

	sort.Strings(domains)

	for _, dom := range domains {
		idx, ok := s.indexReference[dom]
		if !ok {
			// if the domain name is not in the ref map
			// add to the Stats slice and store index in ref map
			s.stats = append(s.stats, DomainStat{
				Name:  dom,
				Count: 1,
			})

			s.indexReference[dom] = len(s.stats) - 1
			continue
		}

		s.stats[idx].Count++
	}

	return nil
}

func (s *Store) GetAllStats() []DomainStat {
	return s.stats
}

func (s *Store) GetDomainCountByName(domainName string) int32 {
	idx, ok := s.indexReference[domainName]
	if !ok {
		return 0
	}

	return s.stats[idx].Count
}
