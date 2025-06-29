package service

import (
	"reflect"
	"strconv"
	"testing"
)

var packSizes = []int{250, 500, 1000, 2000, 5000}

func TestOptimizePacks(t *testing.T) {
	tests := []struct {
		order    int
		expected PackResult
	}{
		{1, PackResult{map[int]int{250: 1}, 250}},
		{250, PackResult{map[int]int{250: 1}, 250}},
		{251, PackResult{map[int]int{500: 1}, 500}},
		{501, PackResult{map[int]int{250: 1, 500: 1}, 750}},
		{12001, PackResult{map[int]int{250: 1, 2000: 1, 5000: 2}, 12250}},
		{0, PackResult{map[int]int{}, 0}},
	}

	for _, test := range tests {
		t.Run("Order_"+strconv.Itoa(test.order), func(t *testing.T) {
			result := OptimizePacks(test.order, packSizes)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("unexpected result for order %d.\nExpected: %+v\nGot:      %+v",
					test.order, test.expected, result)
			}
		})
	}
}
