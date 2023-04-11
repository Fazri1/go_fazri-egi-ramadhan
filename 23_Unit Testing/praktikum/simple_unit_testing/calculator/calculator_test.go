package calculator

import "testing"

func TestAdd(t *testing.T) {
	if Add(10, 2) != 12 {
		t.Error("10 + 2 expected 12")
	}

	if Add(5, 10, 2) != 17 {
		t.Error("5 + 10 + 2 expected 17")
	}

}

func TestSub(t *testing.T) {
	if Sub(10, 2) != 8 {
		t.Error("10 - 2 expected 8")
	}
	if Sub(10, 7) != 3 {
		t.Error("10 - 7 expected 3")
	}

}

func TestMult(t *testing.T) {
	if Mult(10, 2) != 20 {
		t.Error("10 * 2 expected 20")
	}
	if Mult(10, 2, 3) != 60 {
		t.Error("10 * 2 * 3 expected 60")
	}

}

func TestDiv(t *testing.T) {
	if Div(10, 2) != 5 {
		t.Error("10 / 2 expected 5")
	}
	if Div(20, 2, 5) != 2 {
		t.Error("20 / 2 / 5 expected 2")
	}
}
