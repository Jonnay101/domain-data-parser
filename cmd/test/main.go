package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jonnay101/domain-data-parser/pkg/emaildomainstats"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	domainStore, err := timerDecorator(emaildomainstats.GetDomainStats, "customer_data.csv")
	if err != nil {
		return err
	}

	domainData := domainStore.GetAllDomainNameData()

	jsonData, err := json.MarshalIndent(domainData, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}

func timerDecorator(f func(string) (emaildomainstats.DomainStats, error), arg string) (emaildomainstats.DomainStats, error) {
	start := time.Now()
	res, err := f(arg)
	elapsed := time.Since(start)
	log.Printf("ParsEmailsByDomain took %s", elapsed)

	return res, err
}
