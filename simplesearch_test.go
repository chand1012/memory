package simplesearch

import (
	"os"
	"testing"
)

func TestMemoryIndex(t *testing.T) {
	// Define a temporary path for the index file
	testIndexPath := "test.bleve"

	// Create a new memory index
	m, _, err := New(testIndexPath)
	if err != nil {
		t.Fatalf("Failed to create memory index: %v", err)
	}

	// Add memory fragments to the index
	err = m.Add("1", "The quick brown fox jumps over the lazy dog.")
	if err != nil {
		t.Fatalf("Failed to add memory fragment: %v", err)
	}
	err = m.Add("2", "John and Mary went to the market.")
	if err != nil {
		t.Fatalf("Failed to add memory fragment: %v", err)
	}

	// Search the index for a query
	results, err := m.Search("quick brown")
	if err != nil {
		t.Fatalf("Failed to search memory index: %v", err)
	}

	// Check the search results
	if len(results) != 1 || results[0].ID != "1" {
		t.Errorf("Unexpected search results: %+v", results)
	}

	// Test the TopFragmentScore function
	top := TopFragmentScore(results)
	if top.ID != "1" {
		t.Errorf("Unexpected top fragment: %+v", top)
	}

	// Close the memory index
	err = m.Close()
	if err != nil {
		t.Fatalf("Failed to close memory index: %v", err)
	}

	// Destroy the memory index and delete the index file
	err = m.Destroy()
	if err != nil {
		t.Fatalf("Failed to destroy memory index: %v", err)
	}

	// Verify that the index file has been deleted
	if _, err := os.Stat(testIndexPath); !os.IsNotExist(err) {
		t.Errorf("Index file was not deleted: %v", err)
	}
}
