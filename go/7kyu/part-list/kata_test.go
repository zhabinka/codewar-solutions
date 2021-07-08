package kata

import "testing"

func TestPartList(t *testing.T) {
	type args struct {
		list []string
	}
	type test struct {
		name string
		args args
		want string
	}
	tests := []test{
		{
			name: "I, wish, I, hadn't, come",
			args: args{list: []string{"I", "wish", "I", "hadn't", "come"}},
			want: "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)",
		},
		{
			name: "cdIw, tzIy, xDu, rThG",
			args: args{list: []string{"cdIw", "tzIy", "xDu", "rThG"}},
			want: "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartList(tt.args.list); got != tt.want {
				t.Errorf("PartList() -> %s, want %s", got, tt.want)
			}
		})
	}
}
