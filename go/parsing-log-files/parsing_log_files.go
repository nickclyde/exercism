package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	re := regexp.MustCompile(pattern)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	pattern := `<[=~*-]*>`
	re := regexp.MustCompile(pattern)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := `(?i)".*password.*"`
	re := regexp.MustCompile(pattern)
	var count int
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d*`
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User\s+(\S*)\s`
	re := regexp.MustCompile(pattern)
	var taggedLines []string
	for _, line := range lines {
		username := re.FindStringSubmatch(line)
		if username != nil {
			line = fmt.Sprintf("[USR] %s %s", username[1], line)
		}
		taggedLines = append(taggedLines, line)
	}
	return taggedLines
}
