package arithmetic

import (
	"testing"
)

func TestPlus(t *testing.T) {
	var r int

	r = Plus(0, 0)
	if r != 0 {
		t.Errorf("Plus(0, 0) => %d", r)
	}

	r = Plus(42, 10)
	if r != 52 {
		t.Errorf("Plus(42, 10) => %d", r)
	}

	r = Plus(10, 20, 30)
	if r != 60 {
		t.Errorf("Plus(10, 20, 30) => %d", r)
	}

	r = Plus(10, -10)
	if r != 0 {
		t.Errorf("Plus(10, -10) => %d", r)
	}
}

func TestMinus(t *testing.T) {
	var r int

	r = Minus(0, 0)
	if r != 0 {
		t.Errorf("Minus(0, 0) => %d", r)
	}

	r = Minus(42, 10)
	if r != 32 {
		t.Errorf("Minus(42, 0) => %d", r)
	}

	r = Minus(20, 10, 30)
	if r != -20 {
		t.Errorf("Minus(20, 10, 30) => %d", r)
	}

	r = Minus(10, -10)
	if r != 20 {
		t.Errorf("Minus(10, -10) => %d", r)
	}
}

func TestTimes(t *testing.T) {
	var r int

	r = Times(0, 0)
	if r != 0 {
		t.Errorf("Times(0, 0) => %d", r)
	}

	r = Times(42, 10)
	if r != 420 {
		t.Errorf("Times(42, 10) => %d", r)
	}

	r = Times(10, 20, 30)
	if r != 6000 {
		t.Errorf("Times(10, 20, 30) => %d", r)
	}

	r = Times(10, -10)
	if r != -100 {
		t.Errorf("Times(10, -10) => %d", r)
	}
}

func TestDivided(t *testing.T) {
	var r int

	r = Divided(10, 10)
	if r != 1 {
		t.Errorf("Divided(10, 10) => %d", r)
	}

	r = Divided(200, 10)
	if r != 20 {
		t.Errorf("Divided(200, 10) => %d", r)
	}

	r = Divided(100, 10, 5)
	if r != 2 {
		t.Errorf("Divided(100, 10, 5) => %d", r)
	}

	r = Divided(10, -10)
	if r != -1 {
		t.Errorf("Divided(10, -10) => %d", r)
	}
}
