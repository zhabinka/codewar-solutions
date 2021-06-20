package kata

import "testing"

func TestIsDivisible(t *testing.T) {
	type args struct {
		n, x, y int
	}
	type test struct {
		name string
		args args
		want bool
	}
	tests := []test{
		{name: "Equal numbers", args: args{n: 10, x: 10, y: 10}, want: true},
		{name: "n divisible by x and y", args: args{n: 10, x: 5, y: 2}, want: true},
		{name: "n don't divisible by x and y", args: args{n: 10, x: 5, y: 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDivisible(tt.args.n, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsDivisible() = %t, want %t", got, tt.want)
			}
		})
	}
}
