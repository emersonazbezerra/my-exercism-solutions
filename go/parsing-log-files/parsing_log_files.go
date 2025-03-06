package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~*=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`".*"`)
	for _, l := range lines {
		s := re.FindString(l)
		if strings.Contains(strings.ToLower(s), "password") {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	reUser := regexp.MustCompile(`User +[a-zA-Z0-9]+`)
	reName := regexp.MustCompile(`\s+`)

	for i, l := range lines {
		user := reUser.FindString(l)
		if user != "" {
			name := reName.Split(user, -1)
			lines[i] = "[USR] " + name[1] + " " + l
		}
	}
	return lines
}
