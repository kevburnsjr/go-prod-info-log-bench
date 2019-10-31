package main

import (
	"testing"
)

func BenchmarkZapInfo(b *testing.B) {
	test(b, ZapInfo)
}

func BenchmarkZapSugarInfo(b *testing.B) {
	test(b, ZapSugarInfo)
}

func BenchmarkLogrusInfo(b *testing.B) {
	test(b, LogrusInfo)
}

func BenchmarkZapInfoWithFields(b *testing.B) {
	test(b, ZapInfoWithFields)
}

func BenchmarkZapSugarInfoWithFields(b *testing.B) {
	test(b, ZapSugarInfoWithFields)
}

func BenchmarkLogrusInfoWithFields(b *testing.B) {
	test(b, LogrusInfoWithFields)
}

func BenchmarkZapWarn(b *testing.B) {
	test(b, ZapWarn)
}

func BenchmarkZapSugarWarn(b *testing.B) {
	test(b, ZapSugarWarn)
}

func BenchmarkLogrusWarn(b *testing.B) {
	test(b, LogrusWarn)
}

func BenchmarkZapSugarWarnWithFields(b *testing.B) {
	test(b, ZapSugarWarnWithFields)
}

func BenchmarkZapWarnWithFields(b *testing.B) {
	test(b, ZapWarnWithFields)
}

func BenchmarkLogrusWarnWithFields(b *testing.B) {
	test(b, LogrusWarnWithFields)
}

func test(b *testing.B, fn func(string)) {
	for i := 0; i < b.N; i++ {
		fn("test")
	}
}
