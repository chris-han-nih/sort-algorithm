package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1,2,3", args{[]int{2, 1, 3}}, []int{1, 2, 3}},
		{"1,1,1", args{[]int{1, 1, 1}}, []int{1, 1, 1}},
		{"3.5.2.4.1.7.10,6", args{[]int{3, 5, 2, 4, 1, 7, 10, 6}}, []int{1, 2, 3, 4, 5, 6, 7, 10}},
		{"empty", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Sort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Sort() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
