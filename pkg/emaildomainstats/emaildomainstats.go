package emaildomainstats

import "github.com/jonnay101/domain-data-parser/pkg/store"

type DomainStats interface {
	store.DomainData
}

// GetDomainStats returns a DomainData store. This store maps domain names against
// their frequency of occurence in the provided customer data CSV file
func GetDomainStats(filepath string) (DomainStats, error) {
	emailChan := parseCSVData(filepath)
	domChan := parseDomainNames(emailChan)
	return saveDomainStats(domChan)
}
