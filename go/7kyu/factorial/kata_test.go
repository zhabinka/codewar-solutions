package kata

import "testing"

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "0", args: args{n: 0}, want: 1},
		{name: "1", args: args{n: 1}, want: 1},
		{name: "4", args: args{n: 4}, want: 24},
		{name: "7", args: args{n: 7}, want: 5040},
		{name: "17", args: args{n: 17}, want: 355687428096000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() -> %d, want %d", got, tt.want)
			}
		})
	}

}
