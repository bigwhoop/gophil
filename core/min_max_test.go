package core

import (
	"testing"
)

var testData = []int{44,67,1,3,6453,0,12,-645,534}

func TestMinInt(t *testing.T) {
	e := -645
	
	if r := MinInt(testData); e != r {
		t.Errorf("Expected min value to be %d, got %d.", e, r)
	}
}

func TestMaxInt(t *testing.T) {
	e := 6453
	
	if r := MaxInt(testData); e != r {
		t.Errorf("Expected max value to be %d, got %d.", e, r)
	}
}
