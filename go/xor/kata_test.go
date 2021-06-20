package kata

import "testing"

func TestXor(t *testing.T) {
	type args struct {
		a, b bool
	}
	type test struct {
		name string
		args args
		want bool
	}
	tests := []test{
		{name: "false == false", args: args{false, false}, want: false},
		{name: "true == false", args: args{true, false}, want: true},
		{name: "false == true", args: args{false, true}, want: true},
		{name: "true == true", args: args{true, true}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Xor(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Xor() = %t, want %t", got, tt.want)
			}
		})
	}
}
