package datapipeline

import "sync"

func (dp *dataPipeline) mergeStringChannels(chans ...<-chan string) <-chan string {
	merge := make(chan string)
	wg := new(sync.WaitGroup)

	output := func(doms <-chan string) {
		defer wg.Done()

		for dom := range doms {
			merge <- dom
		}
	}

	wg.Add(len(chans))

	for _, ch := range chans {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(merge)
	}()

	return merge
}
