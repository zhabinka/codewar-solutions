package kata

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "strengthlessnesses", args: args{s: "strengthlessnesses"}, want: 1},
		// TODO: {name: "codewarriorsaie", args: args{s: "codewarriorsaie"}, want: 3},
		{name: "iiihoovaeaaaoougjyaw", args: args{s: "yiiihoovaeaaaoougjyaw"}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.s); got != tt.want {
				t.Errorf("Solve() -> %d, want %d", got, tt.want)
			}
		})
	}
}
