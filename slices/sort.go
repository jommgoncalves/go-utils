package slices

import "sort"

// Pair is a data structure to hold key/value pairs.
type Pair struct {
	Key   int
	Value int
}

// PairList is a slice of pairs that implements sort.Interface to sort by values.
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// SortByValues Sorts slice by value.
func SortByValues(data map[int]int) PairList {
	p := make(PairList, len(data))

	i := 0
	for k, v := range data {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	return p
}
