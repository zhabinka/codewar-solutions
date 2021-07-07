package kata

import "testing"

func TestGoals(t *testing.T) {
	type args struct {
		laLigaGoals          int
		copaDelReyGoals      int
		championsLeagueGoals int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Cout goal: 5 + 1 + 8", args: args{laLigaGoals: 5, copaDelReyGoals: 1, championsLeagueGoals: 8}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Goals(tt.args.laLigaGoals, tt.args.copaDelReyGoals, tt.args.championsLeagueGoals); got != tt.want {
				t.Errorf("Goal() = %d, want %d", got, tt.want)
			}
		})
	}
}
