package core

import "testing"

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(RandInt(10))
	}
}
