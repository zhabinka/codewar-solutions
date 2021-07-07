package kata

import "testing"

func TestNthEven(t *testing.T) {
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "1st even number", args: args{n: 1}, want: 0},
		{name: "2nd even number", args: args{n: 2}, want: 2},
		{name: "3th even number", args: args{n: 3}, want: 4},
		{name: "100th even number", args: args{n: 100}, want: 198},
		{name: "1298734th even number", args: args{n: 1298734}, want: 2597466},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NthEven(tt.args.n); got != tt.want {
				t.Errorf("NthEven() = %d, want %d", got, tt.want)
			}
		})
	}
}
