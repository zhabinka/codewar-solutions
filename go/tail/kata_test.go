package kata

import "testing"

func TestCorrectTail(t *testing.T) {
	type args struct {
		body string
		tail rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Fox -> x", args: args{body: "Fox", tail: 'x'}, want: true},
		{name: "Rhino -> x", args: args{body: "Rhino", tail: 'x'}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CorrectTail(tt.args.body, tt.args.tail); got != tt.want {
				t.Errorf("CorrectTail() = %v, want %v", got, tt.want)
			}
		})
	}
}
