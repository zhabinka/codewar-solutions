package kata

import "testing"

func TestRemoveChar(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Remove letter", args: args{word: "golang"}, want: "olan"},
		{name: "Remove space", args: args{word: " golang "}, want: "golang"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveChar(tt.args.word); got != tt.want {
				t.Errorf("RemoveChar() = %q, want %q", got, tt.want)
			}
		})
	}
}
