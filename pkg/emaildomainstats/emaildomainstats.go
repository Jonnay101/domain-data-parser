package emaildomainstats

import (
	"github.com/jonnay101/domain-data-parser/pkg/datapipeline"
	"github.com/jonnay101/domain-data-parser/pkg/store"
)

type DomainStats interface {
	GetAll() map[string]int
	GetByDomainName(domainName string) int
}

func CreateDomainStats(filepath string) (DomainStats, error) {
	dataStore := store.New()
	dataPipeline := datapipeline.New(dataStore)

	if err := dataPipeline.Run(filepath); err != nil {
		return nil, err
	}

	return dataStore, nil
}
