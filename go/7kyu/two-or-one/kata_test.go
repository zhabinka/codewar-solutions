package kata

import (
	"testing"
)

func TestTwoOrOne(t *testing.T) {
	type args struct {
		s1, s2 string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "#1", args: args{s1: "aretheyhere", s2: "yestheyarehere"}, want: "aehrsty"},
		{name: "#2", args: args{s1: "loopingisfunbutdangerous", s2: "lessdangerousthancoding"}, want: "abcdefghilnoprstu"},
	}

	for _, tt := range tests {
		if got := TwoOrOne(tt.args.s1, tt.args.s2); got != tt.want {
			t.Errorf("TwoOrOne() -> %t, want %t", got, tt.want)
		}
	}
}
