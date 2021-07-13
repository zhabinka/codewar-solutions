package kata

import "testing"

func TestFindShort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "String", args: args{s: "bitcoin take over the world maybe who knows perhaps"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShort(tt.args.s); got != tt.want {
				t.Errorf("FindShort() -> %d, want %d", got, tt.want)
			}
		})
	}
}
