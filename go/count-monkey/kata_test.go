package kata

import (
	"reflect"
	"testing"
)

func TestCountMonkeys(t *testing.T) {
	type args struct {
		n int
	}
	type test struct {
		name string
		args args
		want []int
	}
	tests := []test{
		{name: "Count by 1", args: args{n: 1}, want: []int{1}},
		{name: "Count by 10", args: args{n: 10}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountMonkeys(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountMonkeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
