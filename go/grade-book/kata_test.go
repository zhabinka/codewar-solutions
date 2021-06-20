package kata

import "testing"

func TestGetGrade(t *testing.T) {
	type args struct {
		a, b, c int
	}
	type test struct {
		name string
		args args
		want rune
	}
	tests := []test{
		{name: "Letter A", args: args{a: 90, b: 89, c: 91}, want: 'A'},
		{name: "Letter B", args: args{a: 70, b: 82, c: 89}, want: 'B'},
		{name: "Letter C", args: args{a: 70, b: 70, c: 70}, want: 'C'},
		{name: "Letter D", args: args{a: 65, b: 66, c: 58}, want: 'D'},
		{name: "Letter F", args: args{a: 44, b: 48, c: 58}, want: 'F'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGrade(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("GetGrade() = %q, want %q", got, tt.want)
			}
		})
	}
}
