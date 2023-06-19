package searchindex_test

import (
	"reflect"
	"testing"

	searchindex "github.com/stefanovazzocell/SearchIndex"
)

func TestHelpers(t *testing.T) {
	testCases := map[string][]string{
		"  -  ! ":                  {},
		"Hello World!":             {"hello", "world"},
		">  2001: A Space Odyssey": {"2001", "a", "space", "odyssey"},
		"Hello World, Hello User":  {"hello", "world", "user"},
	}
	for test, expected := range testCases {
		tokens := searchindex.GetTokens(test, false)
		if !reflect.DeepEqual(tokens, expected) {
			t.Fatalf("Failed to tokenize %q into %v: got %v", test, expected, tokens)
		}
	}
	// Check sorting
	tokens := searchindex.GetTokens("am I a developer", true)
	if len(tokens) != 4 {
		t.Fatalf("Expected 4 tokens, instead got %v", tokens)
	}
	if tokens[0] != "developer" || tokens[1] != "am" || tokens[2] != "i" || tokens[3] != "a" {
		t.Fatalf("Expected token sorted from longest to shortest, instead got %v", tokens)
	}
}
