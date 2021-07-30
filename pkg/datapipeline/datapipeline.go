package datapipeline

// import "github.com/jonnay101/domain-data-parser/pkg/store"

type DomainStats interface{}

// // GetDomainStats returns a DomainData store. This store maps domain names against
// // their frequency of occurence in the provided customer data CSV file
// func GetDomainStats(filepath string) (DomainStats, error) {
// 	emailChan := parseCSVData(filepath)
// 	domChan := parseDomainNames(emailChan)
// 	return saveDomainStats(domChan)
// }

type DataPipeline interface {
	Run(filepath string) error
}

type dataStore interface {
	SaveDomainName(domainName string) error
}

type dataPipeline struct {
	store dataStore
}

func New(store dataStore) DataPipeline {
	return &dataPipeline{
		store: store,
	}
}

func (dp *dataPipeline) Run(filepath string) error {
	return dp.saveDomainStats(dp.parseDomainNames(dp.parseCSVData(filepath)))
}
