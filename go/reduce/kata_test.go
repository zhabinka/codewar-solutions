package kata

import "testing"

func TestGrow(t *testing.T) {
	type args struct {
		arr []int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "List without zero", args: args{[]int{1, 3, 6, 2}}, want: 36},
		{name: "List with zero", args: args{[]int{1, 3, 0, 6, 2}}, want: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Grow(tt.args.arr); got != tt.want {
				t.Errorf("Grow() -> %d, want %d", got, tt.want)
			}
		})
	}
}
