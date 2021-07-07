package kata

import "testing"

func TestMakeUpperCase(t *testing.T) {
	type args struct {
		str string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "Char to upper", args: args{str: "a"}, want: "A"},
		{name: "Word to upper", args: args{str: "golang"}, want: "GOLANG"},
		{name: "String to upper", args: args{str: "My name is Serge"}, want: "MY NAME IS SERGE"},
		{name: "Complex to upper", args: args{str: "It is friday 13 today."}, want: "IT IS FRIDAY 13 TODAY."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeUpperCase(tt.args.str); got != tt.want {
				t.Errorf("MakeUpperCase() = %q, want %q", got, tt.want)
			}
		})
	}
}
