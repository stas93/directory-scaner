package directory_scaner

import "testing"

var (
	parse = "/home/stas"
)

func BenchmarkScan(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Scan(parse)
	}
}
func BenchmarkScan2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Scan2(parse)
	}
}
func BenchmarkScan3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Scan3(parse)
	}
}
