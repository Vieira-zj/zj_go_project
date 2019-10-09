package gotests

import (
	"testing"
)

func TestCalculation01(t *testing.T) {
	testData := []struct {
		m        int
		n        int
		expected int
	}{
		{1, 1, 2},
		{2, 3, 5},
		{-1, 2, 1},
	}

	for _, unit := range testData {
		if actual := calAdd(unit.m, unit.n); actual != unit.expected {
			t.Errorf("actual: %d, but expected: %d", actual, unit.expected)
		}
	}
}

func TestCalculation02(t *testing.T) {
	testData := []struct {
		m        int
		expected int
	}{
		{1, 2},
		{2, 4},
		{10, 14},
	}

	cal := NewMyCal(1)
	for _, unit := range testData {
		if actual := cal.addAndGet(unit.m); actual != unit.expected {
			t.Errorf("actual: %d, but expected: %d", actual, unit.expected)
		}
	}
}

func TestCalculation03(t *testing.T) {
	var (
		expected = 0
		cal      = NewMyCal(1)
	)
	actual := cal.selfAdd(2).selfAdd(6).selfDivide(8).selfDivide(1).getValue()
	if actual != expected {
		t.Errorf("actual: %d, but expected: %d", actual, expected)
	}
}
