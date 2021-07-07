package kata

import "testing"

func TestCoutPoints(t *testing.T) {
	type args struct {
		games []string
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "All games lost", args: args{games: []string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}}, want: 0},
		{name: "Intense match", args: args{games: []string{"1:1", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"}}, want: 13},
		{name: "All dead heat", args: args{games: []string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountPoints(tt.args.games); got != tt.want {
				t.Errorf("CountPoints() = %d, want %d", got, tt.want)
			}
		})
	}
}

var games = []string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}

func BenchmarkCountPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPoints(games)
	}
}

func BenchmarkCountPoints2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPoints2(games)
	}
}

func BenchmarkCountPoints3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPoints3(games)
	}
}
