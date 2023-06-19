package searchindex

// Searches some text in the index, returns most promising results, if any
func (searchIndex *SearchIndex[Idenfier]) Lookup(text string) []Idenfier {
	// Score idenfiers
	idenfierScore := map[Idenfier]int{}
	bestScore := 0
	for _, token := range GetTokens(text, true) {
		idenfiers, ok := searchIndex.tokensIndex[token]
		if !ok {
			continue
		}
		for _, idenfier := range idenfiers {
			score := idenfierScore[idenfier] + 1
			idenfierScore[idenfier] = score
			if score > bestScore {
				bestScore = score
			}
		}
	}
	// Filter by score
	if len(idenfierScore) == 0 {
		return []Idenfier{} // Found none
	}
	idenfiers := make([]Idenfier, 0, len(idenfierScore))
	for idenfier, score := range idenfierScore {
		if score == bestScore {
			idenfiers = append(idenfiers, idenfier)
		}
	}
	return idenfiers
}
