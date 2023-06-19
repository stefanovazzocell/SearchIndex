package searchindex

import (
	"sort"
	"strings"
	"unicode"
)

// Cleans up strings by removing all non-letters and non-numbers
//
// Optionally sorts by length
func GetTokens(str string, sorted bool) []string {
	tokens := strings.FieldsFunc(strings.ToLower(str), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	seen := make(map[string]bool, len(tokens))
	tokensUniq := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if !seen[token] {
			seen[token] = true
			tokensUniq = append(tokensUniq, token)
		}
	}
	if sorted {
		sort.SliceStable(tokensUniq, func(i, j int) bool {
			return len(tokensUniq[i]) > len(tokensUniq[j])
		})
	}
	return tokensUniq
}
