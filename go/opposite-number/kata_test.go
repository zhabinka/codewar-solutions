package kata

import "testing"

func TestOpposite(t *testing.T) {
	type args struct {
		value int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Positive number", args: args{value: 1}, want: -1},
		{name: "Negative number", args: args{value: -1}, want: 1},
		{name: "Zero number", args: args{value: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Opposite(tt.args.value); got != tt.want {
				t.Errorf("Opposite() = %d, want %d", got, tt.want)
			}
		})
	}
}
