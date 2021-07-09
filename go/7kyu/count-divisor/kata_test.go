package kata

import (
	"testing"
)

func TestDivisors(t *testing.T) {
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Divisors(1)", args: args{n: 1}, want: 1},
		{name: "Divisors(54)", args: args{n: 54}, want: 8},
		{name: "Divisors(64)", args: args{n: 64}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divisors(tt.args.n); got != tt.want {
				t.Errorf("Divisors() -> %d, want %d", got, tt.want)
			}
		})
	}
}
