package watcher

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func ContainKeywords(pattern string, paragraphs ...string) bool {
	for _, paragraph := range paragraphs {
		isMatch, err := regexp.MatchString(pattern, strings.ToLower(paragraph))
		if err != nil {
			log.Fatal(err)
		}
		if isMatch {
			return true
		}
	}
	return false
}

func GenerateRegexp(filters string) string {
	filterArr := strings.Split(filters, ",")
	for i, filter := range filterArr {
		filterArr[i] = fmt.Sprintf("(%s)", strings.TrimSpace(filter))
	}
	return fmt.Sprintf("%s", strings.Join(filterArr, "|"))
}
