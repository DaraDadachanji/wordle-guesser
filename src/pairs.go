package main

import (
	"sort"
)

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (pair PairList) Swap(i, j int)      { pair[i], pair[j] = pair[j], pair[i] }
func (pair PairList) Len() int           { return len(pair) }
func (pair PairList) Less(i, j int) bool { return pair[i].Value < pair[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(_map map[string]int, ascending bool) PairList {
	list := make(PairList, len(_map))
	i := 0
	for key, value := range _map {
		list[i] = Pair{Key: key, Value: value}
		i++
	}
	if ascending {
		sort.Sort(list)
	} else {
		sort.Sort(sort.Reverse(list))
	}
	return list
}
