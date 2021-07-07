package kata

import "testing"

func TestMakeNegative(t *testing.T) {
	type args struct {
		x int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "42", args: args{x: 42}, want: -42},
		{name: "-42", args: args{x: -42}, want: -42},
		{name: "zero", args: args{x: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeNegative(tt.args.x)
			if got != tt.want {
				t.Errorf("MakeNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}
