package kata

import "testing"

func TestHero(t *testing.T) {
	type args struct {
		bullets int
		dragon  int
	}
	type test struct {
		name string
		args args
		want bool
	}
	tests := []test{
		{name: "Enough bullets", args: args{bullets: 5, dragon: 2}, want: true},
		{name: "The hero died", args: args{bullets: 5, dragon: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hero(tt.args.bullets, tt.args.dragon); got != tt.want {
				t.Errorf("Hero() = %t, want %t", got, tt.want)
			}
		})
	}
}
