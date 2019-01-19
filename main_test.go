package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	result := Sum(1, 5)

	if result != 6 {
		t.Errorf("It should be %d and instead got %d", 6, result)
	}

}
