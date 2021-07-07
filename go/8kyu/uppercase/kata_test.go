package kata

import "testing"

func TestMyString_IsUpperCase(t *testing.T) {
	tests := []struct {
		name string
		s    MyString
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsUpperCase(); got != tt.want {
				t.Errorf("IsUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}