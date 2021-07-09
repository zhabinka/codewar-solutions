package kata

import "testing"

func TestHeron(t *testing.T) {
	type args struct {
		a, b, c int
	}
	type test struct {
		name string
		args args
		want float32
	}
	tests := []test{
		{name: "3, 4, 5", args: args{a: 3, b: 4, c: 5}, want: 6.0},
		{name: "4, 5, 3", args: args{a: 4, b: 5, c: 3}, want: 6.0},
		{name: "11, 13, 7", args: args{a: 11, b: 13, c: 7}, want: 38.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Heron(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Heron() -> %f, want %f", got, tt.want)
			}
		})
	}
}
