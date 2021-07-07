package kata

import "testing"

func TestExpressionMatter(t *testing.T) {
	type args struct {
		a, b, c int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Test 1, 1, 1", args: args{a: 1, b: 1, c: 1}, want: 3},
		{name: "Test 3, 5, 7", args: args{a: 3, b: 5, c: 7}, want: 105},
		{name: "Test 2, 10, 3", args: args{a: 2, b: 10, c: 3}, want: 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpressionMatter(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("ExpressionMatter() = %d, want %d", got, tt.want)
			}
		})
	}
}
