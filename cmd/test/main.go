package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jonnay101/domain-data-parser/pkg/emaildomainstats"
	"github.com/jonnay101/domain-data-parser/pkg/persist"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	domainStats, err := timerDecorator(emaildomainstats.GetDomainStats, "customer_data.csv")
	if err != nil {
		return err
	}

	domainData := domainStats.GetAllDomainNameData()

	// persist the data
	db, err := persist.New("./temp")
	if err != nil {
		return err
	}

	if err := db.Save("customer_domain_data.csv", domainData); err != nil {
		return err
	}

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
