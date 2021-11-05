package main
import (
	"regexp"
	"fmt"
)

func RegexMain(s string, p string) bool {
	matched, err := regexp.MatchString("^"+p+"$", s)
	if err != nil {
		return false
	}
	fmt.Println(s, p, matched)
	return matched
}