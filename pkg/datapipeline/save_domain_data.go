package datapipeline

// pipe 3
func (dp *dataPipeline) saveDomainStats(doms <-chan string) error {
	for dom := range doms {
		if err := dp.store.SaveDomainName(dom); err != nil {
			return err
		}
	}

	return nil
}
