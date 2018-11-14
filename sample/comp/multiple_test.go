package comp

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSingleCore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleCore(randomPairs())
	}
}

func BenchmarkNonBufferChanGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NonBufferChanGoRoutine(randomPairs())
	}
}

func BenchmarkBufferChanGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BufferChanGoRoutine(randomPairs())
	}
}

func randomPairs() []MultiplePair {
	var pairs []MultiplePair
	for i := 0; i < 1000; i++ {
		x := random(1, 1000)
		y := random(1, 1000)
		pairs = append(pairs, MultiplePair{x, y})
	}
	return pairs
}

func random(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}
