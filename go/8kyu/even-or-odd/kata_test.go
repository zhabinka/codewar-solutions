package kata

import "testing"

func TestEvenOrOdd(t *testing.T) {
	type args struct {
		number int
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "Odd number", args: args{number: 7}, want: "Odd"},
		{name: "Even number", args: args{number: 8}, want: "Even"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenOrOdd3(tt.args.number); got != tt.want {
				t.Errorf("EvenOrOdd() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkEvenOrOdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvenOrOdd(777)
	}
}

func BenchmarkEvenOrOdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvenOrOdd2(777)
	}
}

func BenchmarkEvenOrOdd3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvenOrOdd3(777)
	}
}
