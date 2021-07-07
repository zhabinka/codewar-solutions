package kata

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		str, ending string
	}
	type test struct {
		name string
		args args
		want bool
	}
	tests := []test{
		{name: "Empty string", args: args{str: "", ending: ""}, want: true},
		{name: "Empty space", args: args{str: " ", ending: ""}, want: true},
		{name: "banana", args: args{str: "banana", ending: "ana"}, want: true},
		{name: "str length < ending length", args: args{str: "", ending: "a"}, want: false},
		{name: "Digits", args: args{str: "1oo", ending: "100"}, want: false},
		{name: "Special symbols", args: args{str: "$a = $b + 1", ending: "+1"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.str, tt.args.ending); got != tt.want {
				t.Errorf("solution() -> %t, want %t", got, tt.want)
			}
		})
	}
}
