package datapipeline

type DomainStats interface{}

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
