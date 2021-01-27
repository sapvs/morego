// Different versions of fizzbuzz implementation for benchmarking
package fizzbuzz

import (
	"testing"
)

var result []interface{}

const COUNT = 50

func BenchmarkFizzBuzzV1(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			result = fizzbuzzV1(COUNT)
		}
	})
}

func BenchmarkFizzBuzzV2(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			result = fizzbuzzV2(COUNT)
		}
	})
}

func BenchmarkFizzBuzzV3(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			result = fizzbuzzV3(COUNT)

		}
	})
}
