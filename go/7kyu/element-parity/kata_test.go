package kata

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		arr []int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Positive int", args: args{arr: []int{1, -1, 2, -2, 3}}, want: 3},
		{name: "Negatve int", args: args{arr: []int{-3, 1, 2, 3, -1, -4, -2}}, want: -4},
		{name: "Double int", args: args{arr: []int{-9, -105, -9, -9, -9, -9, 105}}, want: -9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.arr); got != tt.want {
				t.Errorf("Solve() -> %d, want %d", got, tt.want)
			}
		})
	}
}
