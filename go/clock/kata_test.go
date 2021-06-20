package kata

import "testing"

func TestPast(t *testing.T) {
	type args struct {
		h, m, s int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Zero time", args:args{h: 0, m: 0, s: 0}, want: 0},
		{name: "Zero time", args:args{h: 1, m: 1, s: 1}, want: 3661000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Past(tt.args.h, tt.args.m, tt.args.s); got != tt.want {
				t.Errorf("Past() = %d, want %d", got,tt.want)
			}
		})
	}
}
