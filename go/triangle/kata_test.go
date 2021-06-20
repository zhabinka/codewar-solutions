package kata

import "testing"

func TestOtherAngle(t *testing.T) {
	type args struct {
		a, b int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Test 43, 78", args: args{a: 43, b: 78}, want: 59},
		{name: "Test equilateral triangle", args: args{a: 60, b: 60}, want: 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OtherAngle(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("OtherAngle() = %d, want %d", got, tt.want)
			}
		})
	}
}
