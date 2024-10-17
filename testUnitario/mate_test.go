package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)
	if total != 10 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 10)
	}
}

func TestConvencionalSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, item.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{5, 2, 5},
		{25, 25, 25},
		{-25, 32, 32},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)

		if max != item.c {
			t.Errorf("GetMax() incorrecto, tiene %d se esperaba %d", max, item.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		fibonacci := Fibonacci(item.n)

		if fibonacci != item.r {
			t.Errorf("Funcion Fibonacci() incorrecta, tiene %d se esperaba %d", fibonacci, item.r)
		}
	}
}
