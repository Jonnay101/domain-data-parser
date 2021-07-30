package store

// GetAll returns a map containing the domain names and
// the frequency of their occurence
func (s *domainData) GetAll() map[string]int {
	return s.store
}
