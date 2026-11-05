package wordcount

import (
	"slices"
	"strings"
	"unicode"
)

type Pair struct {
	Word  string
	Count int
}

func Tokens(s string) []string {
	var tokens []string
	var token []rune

	flush := func() {
		if len(token) == 0 {
			return
		}
		tokens = append(tokens, strings.ToLower(string(token)))
		token = token[:0]
	}

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			token = append(token, r)
			continue
		}
		flush()
	}
	flush()

	return tokens
}

// TopN returns up to n pairs, highest count first.
// Ties: alphabetical ascending by Word.
// If n <= 0, return an empty slice.
func TopN(freq map[string]int, n int) []Pair {
	if n <= 0 {
		return []Pair{}
	}

	pairs := make([]Pair, 0, len(freq))
	for word, count := range freq {
		pairs = append(pairs, Pair{Word: word, Count: count})
	}

	slices.SortFunc(pairs, func(a, b Pair) int {
		if a.Count != b.Count {
			if a.Count > b.Count {
				return -1
			}
			return 1
		}
		return strings.Compare(a.Word, b.Word)
	})

	if n < len(pairs) {
		pairs = pairs[:n]
	}
	return pairs
}

// DropStop removes tokens that appear in stop (lowercase keys).
func DropStop(tokens []string, stop map[string]struct{}) []string {
	filtered := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stop[token]; !ok {
			filtered = append(filtered, token)
		}
	}
	return filtered
}
func Count(tokens []string) map[string]int {	
	counts := make(map[string]int)
	for _, token := range tokens {
		counts[token]++
	}
	return counts
}
func TopN(freq map[string]int, n int) []Pair {
	
}
