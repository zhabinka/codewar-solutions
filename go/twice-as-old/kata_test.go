package kata

import "testing"

func TestTwiceAsOld(t *testing.T) {
	type args struct {
		dadYearsOld, sonYearsOld int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Test 42,21", args: args{dadYearsOld: 42, sonYearsOld: 21}, want: 0},
		{name: "Test 55,30", args: args{dadYearsOld: 55, sonYearsOld: 30}, want: 5},
		{name: "Test 36,7", args: args{dadYearsOld: 36, sonYearsOld: 7}, want: 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwiceAsOld(tt.args.dadYearsOld, tt.args.sonYearsOld); got != tt.want {
				t.Errorf("TestTwiceAsOld() = %d, want %d", got, tt.want)
			}
		})
	}
}
