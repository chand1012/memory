package simplesearch

import (
	"os"
	"strings"

	"github.com/blevesearch/bleve"
)

type Memory struct {
	index bleve.Index
	path  string
}

type MemoryFragment struct {
	ID    string
	Score float64
	Avg   float64
}

// New creates a new memory index if one does not exist at the given path.
// If one does exist, it opens it.
func New(path string) (*Memory, bool, error) {
	new := false
	index, err := bleve.Open(path)
	if err != nil {
		new = true
		// create a mapping
		mapping := bleve.NewIndexMapping()
		// create the index
		index, err = bleve.New(path, mapping)
		if err != nil {
			return nil, new, err
		}
	}

	m := &Memory{
		index: index,
		path:  path,
	}

	return m, new, nil
}

// Close closes the index
func (m *Memory) Close() error {
	return m.index.Close()
}

// Destroy closes the index and deletes the index file
func (m *Memory) Destroy() error {
	err := m.index.Close()
	if err != nil {
		return err
	}
	// delete the index file
	return os.RemoveAll(m.path)
}

// Add adds a new memory fragment to the index
func (m *Memory) Add(id string, content string) error {
	return m.index.Index(id, strings.ToLower(content))
}

// Search searches the index for the given query
func (m *Memory) Search(query string) ([]MemoryFragment, error) {
	q := bleve.NewQueryStringQuery(query)
	search := bleve.NewSearchRequest(q)
	searchResults, err := m.index.Search(search)
	if err != nil {
		return nil, err
	}

	// tokenize and search again
	tokens := simpleTokenize(query)
	for _, token := range tokens {
		q = bleve.NewQueryStringQuery(token)
		search = bleve.NewSearchRequest(q)
		r, err := m.index.Search(search)
		if err != nil {
			return nil, err
		}
		searchResults.Hits = append(searchResults.Hits, r.Hits...)
	}

	// Use a map to store the total score and count for each ID
	idScoreCountMap := make(map[string]struct {
		totalScore float64
		count      int
	})

	for _, hit := range searchResults.Hits {
		// Accumulate the score and increment the count for the current ID
		entry := idScoreCountMap[hit.ID]
		entry.totalScore += hit.Score
		entry.count++
		idScoreCountMap[hit.ID] = entry
	}

	var results []MemoryFragment
	for id, entry := range idScoreCountMap {
		// Calculate the average score for each ID
		avg := entry.totalScore / float64(entry.count)

		results = append(results, MemoryFragment{
			ID:    id,
			Score: entry.totalScore,
			Avg:   avg,
		})
	}

	return results, nil
}

// TopFragment returns the fragment with the highest score
func TopFragment(fragments []MemoryFragment) MemoryFragment {
	var top MemoryFragment
	for _, fragment := range fragments {
		if fragment.Score > top.Score {
			top = fragment
		}
	}
	return top
}
