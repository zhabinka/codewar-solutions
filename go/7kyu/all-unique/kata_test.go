package kata

import (
	"testing"
)

func TestHasUniqueChar(t *testing.T) {
	type args struct {
		str string
	}
	type test struct {
		name string
		args args
		want bool
	}
	tests := []test{
		{name: "  nAa", args: args{str: "  nAa"}, want: false},
		{name: "abcde", args: args{str: "abcde"}, want: true},
		{name: "++--", args: args{str: "++--"}, want: false},
		{name: "AaBbC", args: args{str: "AaBbC"}, want: true},
	}

	for _, tt := range tests {
		if got := HasUniqueChar(tt.args.str); got != tt.want {
			t.Errorf("HasUniqueChar() -> %t, want %t", got, tt.want)
		}
	}
}
