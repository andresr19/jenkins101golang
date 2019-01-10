package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	result := Sum(1, 5)

	if result != 10 {
		t.Errorf("It should be %d and instead got %d", 10, 6)
	}

}
