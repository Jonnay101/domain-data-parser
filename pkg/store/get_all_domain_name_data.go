package store

// GetAllDomainNameData returns a map containing the domain names and
// the frequency of their occurence
func (s *domainData) GetAllDomainNameData() map[string]int {
	return s.store
}
