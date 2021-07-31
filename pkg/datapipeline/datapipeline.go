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
	done := make(chan struct{})
	defer close(done)

	emails, errChan := dp.parseCSVData(done, filepath)

	if err := dp.saveDomainStats(
		dp.mergeStringChannels(
			done,
			dp.parseDomainNames(done, emails),
			dp.parseDomainNames(done, emails),
			dp.parseDomainNames(done, emails),
			dp.parseDomainNames(done, emails),
			dp.parseDomainNames(done, emails),
			dp.parseDomainNames(done, emails),
			dp.parseDomainNames(done, emails),
			dp.parseDomainNames(done, emails),
		),
	); err != nil {
		return err
	}

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	default:
	}

	return nil
}
