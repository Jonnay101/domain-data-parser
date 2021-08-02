# Domain Name Data Parser

This packages purpose...
- parses a CSV file
- retreives all user emails from the data
- strips away everything except the email domain names
- counts how many times a domain name appears in the data
- maps this count against the domain name in a store
- allows retreival of all data in the map
- allows the user to find how frequently a domain name occurs in this data set

## In Use

Fetching the domain name data set for your file...
```golang
// fetch the domain stats 
domainStats, err := emaildomainstats.GetDomainNameStats("your_file.csv")
if err != nil {
  return err
}
```

Fetching all the data in map form (marshal for viewing)...
```golang
sortedData, err := domainStats.GetAll() 

byt, err := json.MarshalIndent(dataMap, "", "  ")
if err != nil {
  return err
}

fmt.Printf("%s\n", byt)
```

Query the frequency of a domain name... 
```golang 
domainName := "mirrorweb.com"
domainCount := dataStats.GetDomainCountByName(domainName)


fmt.Printf("The domain name %s appears %d time(s) in this data set", domainName, domainCount)
```