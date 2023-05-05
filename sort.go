package memory

import "sort"

func SortByScore(fragments []MemoryFragment) []MemoryFragment {
	sort.Slice(fragments, func(i, j int) bool {
		return fragments[i].Score > fragments[j].Score
	})
	return fragments
}

func SortByAverage(fragments []MemoryFragment) []MemoryFragment {
	sort.Slice(fragments, func(i, j int) bool {
		return fragments[i].Avg > fragments[j].Avg
	})
	return fragments
}
