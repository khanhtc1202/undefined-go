package main

import "testing"

func BenchmarkDoStub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doStub()
	}
}

func BenchmarkDoStubWithFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doStubWithFmt()
	}
}
