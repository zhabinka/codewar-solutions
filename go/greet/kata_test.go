package kata

import (
	"github.com/brianvoe/gofakeit/v6"
	"testing"
)

func TestGreet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct{
		name string
		args args
		want string
	}{
		{name: "Ryan", args: args{name: "Ryan"}, want: "Hello, Ryan how are you doing today?"},
		{name: "Serge", args: args{name: "Serge"}, want: "Hello, Serge how are you doing today?"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet3(tt.args.name); got != tt.want {
				t.Errorf("Greet() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
	  Greet(gofakeit.Name())
	}
}

func BenchmarkGreet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet2(gofakeit.Name())
	}
}
func BenchmarkGreet3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet3(gofakeit.Name())
	}
}
