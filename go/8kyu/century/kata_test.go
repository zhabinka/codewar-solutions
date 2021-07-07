package kata

import "testing"

func TestCentury(t *testing.T) {
	type args struct {
		year int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "1st century", args: args{year: 89}, want: 1},
		{name: "First year of century", args: args{year: 1901}, want: 20},
		{name: "Last year of century", args: args{year: 2000}, want: 20},
		{name: "Today", args: args{year: 2021}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := century(tt.args.year); got != tt.want {
				t.Errorf("Century() = %d, want %d", got, tt.want)
			}
		})
	}
}
