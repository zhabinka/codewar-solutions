package kata

import "testing"

func TestOddCount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 1", args: args{n: 1}, want: 1},
		{name: "Test 15", args: args{n: 15}, want: 7},
		{name: "Test 15023", args: args{n: 15023}, want: 7511},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OddCount(tt.args.n); got != tt.want {
				t.Errorf("OddCout() = %d, want %d", got, tt.want)
			}
		})
	}
}
