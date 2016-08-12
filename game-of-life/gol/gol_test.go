package gol

import "testing"

func TestGol(t *testing.T) {
	if Gol() == true {
		t.Error()
	}
}
