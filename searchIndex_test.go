package searchindex_test

import (
	"os"
	"reflect"
	"testing"

	searchindex "github.com/stefanovazzocell/SearchIndex"
)

var (
	testData = map[int][]string{
		0: {"Hello world"},
		1: {"Hello user #5!"},
		2: {"lorem", "Laborum quo (voluptatem) accusantium et < nihil error aliquid libero. Saepe quasi eos quasi. Officia explicabo error aut"},
		3: {"Laborum quo accusantium et voluptatem nihil"},
	}
)

func TestSearchIndex(t *testing.T) {
	// New
	index := searchindex.NewSearchIndex[int](testData)
	// Lookup
	if ids := index.Lookup("404"); len(ids) != 0 {
		t.Fatalf("Expected nothing, got %v", ids)
	}
	if ids := index.Lookup("	HeLLO "); !reflect.DeepEqual([]int{0, 1}, ids) {
		t.Fatalf("Expected [0, 1], got %v", ids)
	}
	if ids := index.Lookup("> lorem"); !reflect.DeepEqual([]int{2}, ids) {
		t.Fatalf("Expected [2], got %v", ids)
	}
	if ids := index.Lookup("my favorite number is 5"); !reflect.DeepEqual([]int{1}, ids) {
		t.Fatalf("Expected [1], got %v", ids)
	}
	if ids := index.Lookup("Laborum quo accusantium et voluptatem nihil error"); !reflect.DeepEqual([]int{2}, ids) {
		t.Fatalf("Expected [2], got %v", ids)
	}
	// Save
	os.Remove("./test.index")
	defer os.Remove("./test.index")
	if err := index.Save("./test.index"); err != nil {
		t.Fatalf("Failed to save index: %v", err)
	}
	// Open
	loadedIndex, err := searchindex.OpenSearchIndex[int]("./test.index")
	if err != nil {
		t.Fatalf("Failed to open index: %v", err)
	}
	// Lookup (loadedindex)
	if ids := loadedIndex.Lookup("404"); len(ids) != 0 {
		t.Fatalf("Expected nothing, got %v", ids)
	}
	if ids := loadedIndex.Lookup("	HeLLO "); !reflect.DeepEqual([]int{0, 1}, ids) {
		t.Fatalf("Expected [0, 1], got %v", ids)
	}
	if ids := loadedIndex.Lookup("> lorem"); !reflect.DeepEqual([]int{2}, ids) {
		t.Fatalf("Expected [2], got %v", ids)
	}
	if ids := loadedIndex.Lookup("my favorite number is 5"); !reflect.DeepEqual([]int{1}, ids) {
		t.Fatalf("Expected [1], got %v", ids)
	}
	if ids := loadedIndex.Lookup("Laborum quo accusantium et voluptatem nihil error"); !reflect.DeepEqual([]int{2}, ids) {
		t.Fatalf("Expected [2], got %v", ids)
	}
}
