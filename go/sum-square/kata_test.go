package kata

import "testing"

func TestSumSquare(t *testing.T) {
	type args struct {
		numbers []int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Empty list", args: args{[]int{}}, want: 0},
		{name: "List of numbers", args: args{[]int{0, 3, 4, 5}}, want: 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSquare(tt.args.numbers); got != tt.want {
				t.Errorf("SumSquare() = %d, want %d", got, tt.want)
			}
		})
	}
}
