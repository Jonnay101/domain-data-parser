package store

// GetByDomainName returns the amount of times the domain name was
// found in the customer data CSV file
func (s *domainData) GetByDomainName(domainName string) int {
	freq, ok := s.store[domainName]
	if !ok {
		return 0
	}

	return freq
}
