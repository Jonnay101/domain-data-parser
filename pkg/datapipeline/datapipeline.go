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
	emails := dp.parseCSVData(filepath)

	return dp.saveDomainStats(
		dp.mergeStringChannels(
			dp.parseDomainNames(emails),
			dp.parseDomainNames(emails),
			dp.parseDomainNames(emails),
			dp.parseDomainNames(emails),
			dp.parseDomainNames(emails),
			dp.parseDomainNames(emails),
			dp.parseDomainNames(emails),
			dp.parseDomainNames(emails),
		),
	)
}
