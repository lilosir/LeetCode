package problem811

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domainsMap := make(map[string]int)

	for _, domain := range cpdomains {
		times, _ := strconv.Atoi(strings.Split(domain, " ")[0])
		currentDomain := strings.Split(domain, " ")[1]
		subDomainArr := strings.Split(currentDomain, ".")

		for i := 1; i <= len(subDomainArr); i++ {
			len := len(strings.SplitN(currentDomain, ".", i))
			sub := strings.SplitN(currentDomain, ".", i)[len-1]
			domainsMap[sub] += times
		}
	}

	result := []string{}
	for k, v := range domainsMap {
		times := strconv.Itoa(v)
		result = append(result, times+" "+k)
	}
	return result
}
