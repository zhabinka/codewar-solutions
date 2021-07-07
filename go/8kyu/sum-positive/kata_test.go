package kata

import "testing"

func TestPositiveSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Empty list", args: args{numbers: []int{}}, want: 0},
		{name: "Numbers positive list", args: args{numbers: []int{1, 2, 3, 4}}, want: 10},
		{name: "Numbers different list", args: args{numbers: []int{1, 2, -5, 3, 4, 0}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PositiveSum(tt.args.numbers); got != tt.want {
				t.Errorf("PositiveSum() = %d, want %d", got, tt.want)
			}
		})
	}
}
