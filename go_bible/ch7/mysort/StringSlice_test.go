package mysort_test

import (
	"sort"
	"testing"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestMySort(t *testing.T) {
	names := []string{"a", "e", "d", "f"}
	sort.Sort(StringSlice(names))
}
