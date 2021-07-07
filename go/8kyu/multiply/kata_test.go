package kata

import "testing"

func TestMuliply(t *testing.T) {
	type args struct {
		a, b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Zero", args: args{a: 1, b: 0}, want: 0},
		{name: "5 x 6", args: args{a: 5, b: 6}, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiply() = %d, want %d", got, tt.want)
			}
		})
	}
}
