package echo

import "testing"

func benchmarkEcho(b *testing.B, size int, echo func([]string)) {
	b.StopTimer()
	args := make([]string, size)
	for i := 0; i < size; i++ {
		args[i] = "abcdefg"
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		echo(args)
	}
}

func BenchmarkEcho1_10(b *testing.B) {
	benchmarkEcho(b, 10, echo1)
}
func BenchmarkEcho1_100(b *testing.B) {
	benchmarkEcho(b, 100, echo1)
}
func BenchmarkEcho1_1000(b *testing.B) {
	benchmarkEcho(b, 1000, echo1)
}

func BenchmarkEcho2_10(b *testing.B) {
	benchmarkEcho(b, 10, echo2)
}
func BenchmarkEcho2_100(b *testing.B) {
	benchmarkEcho(b, 100, echo2)
}
func BenchmarkEcho2_1000(b *testing.B) {
	benchmarkEcho(b, 1000, echo2)
}
