package kata

import "testing"

func TestQuarterOf(t *testing.T) {
	type args struct {
		month int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "First quarter", args: args{month: 3}, want: 1},
		{name: "Second quarter", args: args{month: 5}, want: 2},
		{name: "Third quarter", args: args{month: 7}, want: 3},
		{name: "Fourth quarter", args: args{month: 12}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuarterOf(tt.args.month); got != tt.want {
				t.Errorf("Quarter() = %d, want %d", got, tt.want)
			}
		})
	}
}
