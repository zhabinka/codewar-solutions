package kata

import "testing"

func TestPrinterError(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "Correct colors", args: args{s: "aaabbbbhaijjjm"}, want: "0/14"},
		{name: "Incorrect colors", args: args{s: "nopzqrw"}, want: "7/7"},
		{name: "Diff colors", args: args{s: "aaaxbbbbyyhwawiwjjjwwm"}, want: "8/22"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrinterError(tt.args.s); got != tt.want {
				t.Errorf("PrinterError() -> %s, want %s", got, tt.want)
			}
		})
	}
}
