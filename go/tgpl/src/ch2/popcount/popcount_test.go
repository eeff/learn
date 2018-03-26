package popcount

import (
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

var tests = []struct {
	u    uint64
	want int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{7, 3},
	{15, 4},
	{254, 7},
}

func testPopCount(t *testing.T, f func(uint64) int) {
	fullname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fname := strings.TrimPrefix(filepath.Ext(fullname), ".")
	for _, test := range tests {
		if got := f(test.u); got != test.want {
			t.Errorf("%s(%0b) = %#v, want %#v\n", fname, test.u, got, test.want)
		}
	}
}
func TestPopCount(t *testing.T) {
	testPopCount(t, PopCount)
}

func TestPopCountShifting(t *testing.T) {
	testPopCount(t, PopCountShifting)
}

func TestPopCountBitop(t *testing.T) {
	testPopCount(t, PopCountBitop)
}

func benchmarkPopCount(b *testing.B, f func(uint64) int) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	for i := 0; i < b.N; i++ {
		f(rand.Uint64())
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchmarkPopCount(b, PopCount)
}

func BenchmarkPopCountShifting(b *testing.B) {
	benchmarkPopCount(b, PopCountShifting)
}

func BenchmarkPopCountBitop(b *testing.B) {
	benchmarkPopCount(b, PopCountBitop)
}
