package kata

import "testing"

func TestAbrevName(t *testing.T) {
	type args struct {
		name string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "Full name", args: args{name: "Adam Smith"}, want: "A.S"},
		{name: "Short name", args: args{name: "A Smith"}, want: "A.S"},
		{name: "Lowercase name", args: args{name: "adam smith"}, want: "A.S"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbbrevName(tt.args.name); got != tt.want {
				t.Errorf("AbbrevName() = %q, want %q", got, tt.want)
			}
		})
	}
}
