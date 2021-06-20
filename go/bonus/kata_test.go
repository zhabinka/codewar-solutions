package kata

import "testing"

func TestBonusTime(t *testing.T) {
	type args struct {
		salary int
		bonus  bool
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{name: "Without bonus", args: args{salary: 100, bonus: false}, want: "£100"},
		{name: "With bonus", args: args{salary: 100, bonus: true}, want: "£1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BonusTime(tt.args.salary, tt.args.bonus); got != tt.want{
				t.Errorf("BonusTime() = %s want %s", got, tt.want)
			}
		})
	}
}
