package src

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	"1",
		//	args{
		//		[]int{9, 3, 30, 34, 5},
		//	},
		//	"9534330",
		//},
		//{
		//	"2",
		//	args{
		//		[]int{9,900,990,909,54,2,6432,623,4,6,27,2,7,45,23345},
		//	},
		//	"999090990076643262354454272334522",
		//},
		{
			"3",
			args{
				[]int{0,0,0},
			},
			"0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
