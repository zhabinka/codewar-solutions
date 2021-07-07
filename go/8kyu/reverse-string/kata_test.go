package kata

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		str string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "Empty", args: args{str: ""}, want: ""},
		{name: "Char", args: args{str: "a"}, want: "a"},
		{name: "Word", args: args{str: "golang"}, want: "gnalog"},
		{name: "String", args: args{str: "go is go!"}, want: "!og si og"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.str); got != tt.want {
				t.Errorf("Reverse() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("GolangGolangGolangGolangGolang")
	}
}
func BenchmarkReverse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse2("GolangGolangGolangGolangGolang")
	}
}
