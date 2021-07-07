package kata

import "testing"

func TestHighAndLow(t *testing.T) {
	type args struct {
		in string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "8 3 -5 42 -1 0 0 -9 4 7 4 -4", args: args{in: "8 3 -5 42 -1 0 0 -9 4 7 4 -4"}, want: "42 -9"},
		{name: "8 3 5 42 1 9 4 7 4 45", args: args{in: "8 3 5 42 1 9 4 7 4 45"}, want: "45 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighAndLow(tt.args.in); got != tt.want {
				t.Errorf("HighAndLow() -> %q, want %q", got, tt.want)
			}
		})
	}

}
