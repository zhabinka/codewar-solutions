package kata

import "testing"

func TestSummation(t *testing.T) {
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Sum elements up to 1", args: args{n: 1}, want: 1},
		{name: "Sum elements up to 8", args: args{n: 8}, want: 36},
		{name: "Sum elements up to 213", args: args{n: 213}, want: 22791},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Summation(tt.args.n); got != tt.want {
				t.Errorf("Summation() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkSummation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Summation(1000)
	}
}

func BenchmarkSummation2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Summation2(1000)
	}
}

func BenchmarkSummation3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Summation3(1000)
	}
}
