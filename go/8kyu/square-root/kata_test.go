package kata

import (
	"reflect"
	"testing"
)

func TestSquareOrSquareRoot(t *testing.T) {
	type args struct {
		arr []int
	}
	type test struct {
		name string
		args args
		want []int
	}
	tests := []test{
		{name:" Test: {100, 101, 5, 5, 1, 1}", args:args{arr: []int{100, 101, 5, 5, 1, 1}}, want: []int{10, 10201, 25, 25, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareOrSquareRoot2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareOrSquareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSquareOrSquareRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareOrSquareRoot([]int{100, 101, 5, 5, 1, 1})
	}
}

func BenchmarkSquareOrSquareRoot2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareOrSquareRoot2([]int{100, 101, 5, 5, 1, 1})
	}
}
