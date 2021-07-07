package kata

import "testing"

func TestSeatsItTheater(t *testing.T) {
	type args struct {
		nCols, nRows, col, row int
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{
		{name: "Test 1, 1, 1, 1", args: args{nCols: 1, nRows: 1, col: 1, row: 1}, want: 0},
		{name: "Test 16, 11, 5, 3", args: args{nCols: 16, nRows: 11, col: 5, row: 3}, want: 96},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := seatsInTheater(tt.args.nCols, tt.args.nRows, tt.args.col, tt.args.row); got != tt.want {
				t.Errorf("seatsInTheater() = %d, want %d", got, tt.want)
			}
		})
	}
}
