package kata

import (
	"reflect"
	"testing"
)

func TestMultipleOfIndex(t *testing.T) {
	type args struct {
		numbers []int
	}
	type test struct {
		name string
		args args
		want []int
	}
	tests := []test{
		{name: "Empty list", args: args{[]int{}}, want: []int{}},
		{name: "List", args: args{[]int{22, -6, 32, 82, 9, 25}}, want: []int{-6, 32, 25}},
		{name: "List with zero", args: args{[]int{-5, -85, 72, -26, -14, 76, -7, 72, 0}}, want: []int{-85, 72, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multipleOfIndex(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyOfIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
