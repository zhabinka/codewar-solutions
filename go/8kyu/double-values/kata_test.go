package kata

import (
	"reflect"
	"testing"
)

func TestDoubleValues(t *testing.T) {
	type args struct {
		numbers []int
	}
	type test struct {
		name string
		args args
		want []int
	}
	tests := []test{
		{name: "Empty list", args: args{numbers: []int{}}, want: []int{}},
		{name: "Test list", args: args{numbers: []int{0, 2, 6, 7}}, want: []int{0, 4, 12, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleValues(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
