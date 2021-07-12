package kata

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "Without space", args: args{s: "codewars"}, want: "srawedoc"},
		{name: "With middle space", args: args{s: "your code"}, want: "edoc ruoy"},
		{name: "With assimetrical spases", args: args{s: "your code rocks"}, want: "skco redo cruoy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.s); got != tt.want {
				t.Errorf("Solve() -> %s, want %s", got, tt.want)
			}
		})
	}
}
