package dod

import (
	"testing"
)

const (
	maxLen = 5_000_000_0
)

var result int64

func BenchmarkI1(b *testing.B) {
	s := make([]I1, maxLen)
	var r int64
	b.ResetTimer()
	for j := 0; j < maxLen; j++ {
		r += s[j].i
	}
	result = r
}

func BenchmarkI2(b *testing.B) {
	s := make([]I2, maxLen)
	var r int64
	b.ResetTimer()
	for j := 0; j < maxLen; j++ {
		r += s[j].i
	}
	result = r
}
