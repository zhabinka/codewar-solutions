package kata

import "testing"

func TestEasyLine(t *testing.T) {
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want string
	}

	tests := []test{
		{name: "EasyLine(0) – first line", args: args{n: 0}, want: "1"},
		{name: "EasyLine(1)", args: args{n: 1}, want: "2"},
		{name: "EasyLine(7)", args: args{n: 7}, want: "3432"},
		{name: "EasyLine(100)", args: args{n: 100}, want: "90548514656103281165404177077484163874504589675413336841320"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EasyLine(tt.args.n); got != tt.want {
				t.Errorf("EasyLine() -> %s, want %s", got, tt.want)
			}
		})
	}
}
