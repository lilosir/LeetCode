package problem929

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	res := make(map[string]int)
	for _, email := range emails {
		e := validEmail(email)
		res[e]++
	}
	return len(res)
}

func validEmail(email string) string {
	subs := strings.Split(email, "@")
	valid := strings.Split(subs[0], "+")[0]
	local := strings.Join(strings.Split(valid, "."), "")
	return local + "@" + subs[1]
}
