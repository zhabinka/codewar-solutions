package kata

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	type test struct {
		name string
		args args
		want bool
	}
	tests := []test{
		{name: "Empty string", args: args{str: ""}, want: true},
		{name: "String is palindrome", args: args{str: "abba"}, want: true},
		{name: "String isn't palindrome", args: args{str: "hello"}, want: false},
		{name: "Palindrome with uppercase", args: args{str: "Abba"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.str); got != tt.want {
				t.Errorf("IsPalindrome() = %t, want %t", got, tt.want)
			}
		})
	}
}
