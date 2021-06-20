package kata

import "testing"

func TestDNAtoRNA(t *testing.T) {
	type args struct {
		dna string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "Convert GCAT", args: args{dna: "GCAT"}, want: "GCAU"},
		{name: "Convert CTTG", args: args{dna: "CTTG"}, want: "CUUG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DNAtoRNA(tt.args.dna); got != tt.want {
				t.Errorf("DNAtoRNA() = %q, want %q", got, tt.want)
			}
		})
	}
}
