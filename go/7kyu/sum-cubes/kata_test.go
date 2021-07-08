package kata

import "testing"

func TestSumCubes(t *testing.T) {
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "1", args: args{n: 1}, want: 1},
		{name: "2", args: args{n: 2}, want: 9},
		{name: "123", args: args{n: 123}, want: 58155876},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumCubes(tt.args.n); got != tt.want {
				t.Errorf("SumCube() -> %d, want %d", got, tt.want)
			}
		})
	}
}
