package emaildomainstats

import (
	"github.com/jonnay101/domain-data-parser/pkg/store"
)

func saveDomainStats(doms <-chan string) (DomainStats, error) {
	dns := store.New()
	for dom := range doms {
		if err := dns.SaveDomainName(dom); err != nil {
			return nil, err
		}
	}

	return dns, nil
}
