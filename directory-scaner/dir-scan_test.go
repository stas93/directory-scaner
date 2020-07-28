package directory_scaner

import "testing"

var (
	parse = "/home/stas"
)

func BenchmarkScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Scan(parse)
	}
}
func BenchmarkScan2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Scan3(parse)
	}
}
