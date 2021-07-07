package kata

import "testing"

func TestBoolToWord(t *testing.T) {
	type args struct {
		boolean bool
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "True", args: args{boolean: true}, want: "Yes"},
		{name: "False", args: args{boolean: false}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolToWord(tt.args.boolean); got != tt.want {
				t.Errorf("BoolToWord() = %q, want %q", got, tt.want)
			}
		})
	}
}
