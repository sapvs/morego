package sum

import "testing"

type test struct {
	tname string
	in    int
	out   int
}

var tests []*test

func init() {
	tests = []*test{
		{"1", 1, 1},
		{"6", 6, 12},
		{"9999", 9999, 15912},
	}
}

func Test_divisorSumv1(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.tname, func(t *testing.T) {
			if got := divisorSumV1(tt.in); got != tt.out {
				t.Errorf("divisorSumv1() = %v, want %v", got, tt.out)
			}
		})
	}
}

func Test_divisorSumv2(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.tname, func(t *testing.T) {
			if got := divisorSumV2(tt.in); got != tt.out {
				t.Errorf("divisorSumv2() = %v, want %v", got, tt.out)
			}
		})
	}
}

var benchResult int

func BenchmarkDivisorSumv1(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = divisorSumV1(9999)
	}
	benchResult = r
}

func BenchmarkDivisorSumv2(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = divisorSumV2(9999)
	}
	benchResult = r
}
