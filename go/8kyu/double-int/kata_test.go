package kata

import "testing"

func TestDoubleInteger(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Zero", args: args{i: 0}, want: 0},
		{name: "One", args: args{i: 1}, want: 2},
		{name: "Integer", args: args{i: 11}, want: 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleInteger(tt.args.i); got != tt.want {
				t.Errorf("DoubleInteger() = %d, want %d", got, tt.want)
			}
		})
	}
}
