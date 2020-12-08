package util

import (
	"reflect"
	"testing"
)

func Test_closure(t *testing.T) {
	tests := []struct {
		name string
		want func() int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closure() = %v, want %v", got, tt.want)
			}
		})
	}
}
