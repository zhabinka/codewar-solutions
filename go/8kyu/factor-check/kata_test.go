package kata

import "testing"

func TestCheckForFactor(t *testing.T) {
	type args struct {
		base, factor int
	}
	type test struct {
		name string
		args args
		want bool
	}
	tests := []test{
		{name: " 10 % 2 == 0", args: args{base: 10, factor: 2}, want: true},
		{name: " 63 % 7 == 0", args: args{base: 63, factor: 7}, want: true},
		{name: " 11 % 2 == 0", args: args{base: 11, factor: 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckForFactor(tt.args.base, tt.args.factor); got != tt.want {
				t.Errorf("CheckForFactor() = %t, want %t", got, tt.want)
			}
		})
	}
}
