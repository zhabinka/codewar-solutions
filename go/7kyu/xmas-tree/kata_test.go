package kata

import (
	"reflect"
	"testing"
)

func TestXMasTree(t *testing.T) {
	type args struct {
		height int
	}
	type test struct {
		name string
		args args
		want []string
	}
	tests := []test{
		{name: "Tree height 0", args: args{height: 0}, want: []string{}},
		{name: "Tree height 2", args: args{height: 2}, want: []string{"_#_", "###", "_#_", "_#_"}},
		{name: "Tree height 4", args: args{height: 4}, want: []string{"___#___", "__###__", "_#####_", "#######", "___#___", "___#___"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XMasTree(tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMashTree() -> %#v, want %#v", got, tt.want)
			}
		})
	}
}
