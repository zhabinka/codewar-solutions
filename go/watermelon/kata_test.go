package kata

import "testing"

func TestDivideWatermelons(t *testing.T) {
	type args struct {
		amount int
	}
	type test struct {
		name string
		args args
		want bool
	}
	tests := []test{
		{name: "2 watermelons", args: args{amount: 2}, want: false},
		{name: "8 watermelons", args: args{amount: 8}, want: true},
		{name: "13 watermelons", args: args{amount: 13}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DivideWatermelons(tt.args.amount); got != tt.want {
				t.Errorf("DivideMelons() = %t, want %t", got, tt.want)
			}
		})
	}
}
