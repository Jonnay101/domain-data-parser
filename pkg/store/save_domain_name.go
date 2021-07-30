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
		s.store[domainName]++
		return nil
	}

	s.store[domainName] = 1
	return nil
}

func (s *domainData) domainNameExists(domainName string) bool {
	_, ok := s.store[domainName]
	return ok
}
