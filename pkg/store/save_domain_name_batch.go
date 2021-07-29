package store

import (
	"errors"
)

// SaveDomainName adds the domain name to the store. If the name already exists,
// it increments the frequency counter for that domain name.
func (s *domainData) SaveDomainName(domainName string) error {
	return s.saveDomainName(domainName)
}

func (s *domainData) saveDomainName(domainName string) error {
	if domainName == "" {
		return errors.New("invalid domain name")
	}

	if s.domainNameExists(domainName) {
		s.incrementDomainFrequency(domainName)
		return nil
	}

	s.initializeDomainName(domainName)
	return nil
}

func (s *domainData) domainNameExists(domainName string) bool {
	_, ok := s.store[domainName]
	return ok
}

func (s *domainData) incrementDomainFrequency(domainName string) {
	s.store[domainName]++
}

func (s *domainData) initializeDomainName(domainName string) {
	s.store[domainName] = 1
}
