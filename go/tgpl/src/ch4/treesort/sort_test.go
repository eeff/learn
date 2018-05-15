package treesort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	t.Logf("TestSort seed: %d", seed)
	for i := 0; i < 20; i++ {
		data := make([]int, 50)
		for i := range data {
			data[i] = rand.Int()
		}
		got := make([]int, len(data))
		copy(got, data)
		if Sort(got); !sort.IntsAreSorted(got) {
			t.Errorf("Sort(%v) == %v, not sorted", data, got)
		}
	}
}
