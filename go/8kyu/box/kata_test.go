package kata

import "testing"

func TestGetSize(t *testing.T) {
	type args struct {
		w, h, d int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{name: "Small box", args: args{w: 1, h: 1, d: 1}, want: [2]int{6, 1}},
		{name: "Box 4, 2, 6", args: args{w: 4, h: 2, d: 6}, want: [2]int{88, 48}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSize(tt.args.w, tt.args.h, tt.args.d); got != tt.want {
				t.Errorf("GetSize() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
