package searchindex

import (
	"encoding/json"
	"os"
)

// A simple text search index
type SearchIndex[Identifier comparable] struct {
	tokensIndex map[string][]Identifier // Index mapping a token to matching idenfiers
}

// Builds a new text search index
func NewSearchIndex[Identifier comparable](entries map[Identifier][]string) *SearchIndex[Identifier] {
	searchIndex := SearchIndex[Identifier]{
		tokensIndex: map[string][]Identifier{},
	}
	// Build tokens index
	for idenfier, texts := range entries {
		// Gather tokens
		tokensMap := map[string]bool{}
		for _, text := range texts {
			for _, token := range GetTokens(text, false) {
				tokensMap[token] = true
			}
		}
		// Store tokens
		for token := range tokensMap {
			searchIndex.tokensIndex[token] = append(searchIndex.tokensIndex[token], idenfier)
		}
	}
	return &searchIndex
}

func OpenSearchIndex[Identifier comparable](fileName string) (*SearchIndex[Identifier], error) {
	searchIndex := SearchIndex[Identifier]{
		tokensIndex: map[string][]Identifier{},
	}
	// Open file
	file, err := os.Open(fileName) //#nosec
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()
	// Load contents
	err = json.NewDecoder(file).Decode(&searchIndex.tokensIndex)
	return &searchIndex, err
}
