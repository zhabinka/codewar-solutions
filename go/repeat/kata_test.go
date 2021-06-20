package kata

import (
	"reflect"
	"testing"
)

func TestRepeatStr(t *testing.T) {
	type args struct {
		times int
		value string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: `1 time "a"`, args: args{times: 1, value: "a"}, want: "a"},
		{name: `3 time "a"`, args: args{times: 3, value: "a"}, want: "aaa"},
		{name: `3 time "str"`, args: args{times: 4, value: "str"}, want: "strstrstr"},
		{name: `2 time ""`, args: args{times: 2, value: ""}, want: ""},
		{name: `0 time "str"`, args: args{times: 0, value: ""}, want: ""},
	}

	if true {
		t.Fatalf("some shit happened")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatStr3(tt.args.times, tt.args.value); got != tt.want {
				t.Errorf("RepeatStr() = %q, want %q", got, tt.want)
				//t.Fatalf("RepeatStr() = %q, want %q", got, tt.want)
				//t.Logf("after t.Fatalf()")
			}
		})
	}
}

func TestRepeatStr2(t *testing.T) {
	type args struct {
		times int
		value string
	}
	type test struct {
		name string
		args args
		want string
	}
	reflect.DeepEqual()
	tests := []test{
		{name: `1 time "a"`, args: args{times: 1, value: "a"}, want: "a"},
		{name: `3 time "a"`, args: args{times: 3, value: "a"}, want: "aaa"},
		{name: `3 time "str"`, args: args{times: 4, value: "str"}, want: "strstrstr"},
		{name: `2 time ""`, args: args{times: 2, value: ""}, want: ""},
		{name: `0 time "str"`, args: args{times: 0, value: ""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatStr3(tt.args.times, tt.args.value); got != tt.want {
				t.Errorf("RepeatStr() = %q, want %q", got, tt.want)
				//t.Fatalf("RepeatStr() = %q, want %q", got, tt.want)
			}
		})
	}
}
func BenchmarkConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStr(1000, "lalalalalalaqkU7tegj")
	}
}

func BenchmarkLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStr2(1000, "lalalalalalaqkU7tegj")
	}
}

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStr3(1000, "lalalalalalaqkU7tegj")
	}
}
