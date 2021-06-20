package kata

import "testing"

func TestCountSheep(t *testing.T) {
	type args struct {
		num int
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "0 sheep", args: args{num: 0}, want: ""},
		{name: "1 sheep", args: args{num: 1}, want: "1 sheep..."},
		{name: "3 sheep", args: args{num: 3}, want: "1 sheep...2 sheep...3 sheep..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSheep(tt.args.num); got != tt.want {
				t.Errorf("counSheep() = %q, want %q", got, tt.want)
			}
		})
	}
}
