package kata

import "testing"

func TestCombat(t *testing.T) {
	type args struct {
		health, damage float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "health > damage", args: args{health: 100.0, damage: 50.0}, want: 50.0},
		{name: "health < damage", args: args{health: 1.0, damage: 50}, want: 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combat(tt.args.health, tt.args.damage); got != tt.want {
				t.Errorf("combat() = %.1f, want %.1f", got, tt.want)
			}
		})
	}
}
