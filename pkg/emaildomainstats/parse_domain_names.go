package emaildomainstats

import (
	"net/mail"
	"strings"
)

func parseDomainNames(emailChan <-chan string) <-chan string {
	domainChan := make(chan string)

	go func() {
		for email := range emailChan {
			if !emailAddressIsValid(email) {
				continue
			}

			domainChan <- stripDomainName(email)
		}

		close(domainChan)
	}()

	return domainChan
}

func emailAddressIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func stripDomainName(email string) string {
	return strings.Split(email, "@")[1]
}
