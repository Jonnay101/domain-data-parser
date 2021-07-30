package store

type DomainData interface {
	SaveDomainName(domainName string) error
	GetAll() map[string]int
	GetByDomainName(domainName string) int
}

type domainData struct {
	store map[string]int
}

// New creates a new domain data store
func New() DomainData {
	return &domainData{
		store: make(map[string]int),
	}
}
