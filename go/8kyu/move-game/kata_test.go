package kata

import "testing"

func TestMove(t *testing.T) {
	type args struct {
		position, roll int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Start from 0", args: args{position: 0, roll: 4}, want: 8},
		{name: "Start from non-zero position", args: args{position: 3, roll: 6}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Move(tt.args.position, tt.args.roll); got != tt.want {
				t.Errorf("Move() = %d, want %d", got, tt.want)
			}
		})
	}
}
