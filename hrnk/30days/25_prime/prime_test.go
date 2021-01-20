package main

import (
	"testing"
)

var result bool

const iter int = 9999

// Benchmarks
func benchIsPrimeAll(funcToTest func(n int) bool, input int, b *testing.B) {
	var r bool

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 1; i < input; i++ {
				r = funcToTest(i)
			}
		}
	})

	result = r
}

func BenchmarkIsPrime1(b *testing.B) {
	benchIsPrimeAll(isPrimev1, iter, b)
}

func BenchmarkIsPrime2(b *testing.B) {
	benchIsPrimeAll(isPrimev2, iter, b)
}

func BenchmarkIsPrime3(b *testing.B) {
	benchIsPrimeAll(isPrimev3, iter, b)
}

func TestIsPrime1(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{num: 1}, false},
		{"2", args{num: 2}, true},
		{"3", args{num: 3}, true},
		{"4", args{num: 4}, false},
		{"15", args{num: 15}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrimev1(tt.args.num); got != tt.want {
				t.Errorf("IsPrime1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{num: 1}, false},
		{"2", args{num: 2}, true},
		{"3", args{num: 3}, true},
		{"4", args{num: 4}, false},
		{"15", args{num: 15}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrimev2(tt.args.num); got != tt.want {
				t.Errorf("IsPrime2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime3(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{num: 1}, false},
		{"2", args{num: 2}, true},
		{"3", args{num: 3}, true},
		{"4", args{num: 4}, false},
		{"15", args{num: 15}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrimev3(tt.args.num); got != tt.want {
				t.Errorf("IsPrime3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{num: 1}, false},
		{"2", args{num: 2}, true},
		{"3", args{num: 3}, true},
		{"4", args{num: 4}, false},
		{"15", args{num: 15}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.num); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
