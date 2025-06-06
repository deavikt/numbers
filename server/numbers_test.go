package server

import "testing"

func TestAddNumber(t *testing.T) {
	var numbersSizeBefore, numbersSizeAfter int

	srv := Server{
		Numbers:  nil,
		Port:     ":8081",
		DataPath: "/data",
	}
	srv.Numbers = append(srv.Numbers, 1, 2, 3, 4, 5)
	numbersSizeBefore = len(srv.Numbers)

	srv.addNumber(100)
	numbersSizeAfter = len(srv.Numbers)

	if numbersSizeBefore == numbersSizeAfter {
		t.Errorf(
			"Numbers amount hasn't increased, got: %d, want: %d",
			numbersSizeAfter,
			numbersSizeAfter+1)
	}
}

func TestDeleteNumbers(t *testing.T) {
	srv := Server{
		Numbers:  nil,
		Port:     ":8081",
		DataPath: "/data",
	}
	srv.Numbers = append(srv.Numbers, 1, 2, 3, 4, 5)

	srv.deleteNumbers()

	if len(srv.Numbers) != 0 {
		t.Errorf("Numbers amount hasn't decreased, got: %d, want: %d", len(srv.Numbers), 0)
	}
}

func TestGetNumbersSum(t *testing.T) {
	srv := Server{
		Numbers:  nil,
		Port:     ":8081",
		DataPath: "/data",
	}

	tests := []struct {
		a, b, c int
		want    int
	}{
		{1, 2, 3, 6},
		{-3, -2, -1, -6},
		{0, 0, 0, 0},
		{-3, 2, 0, -1},
	}

	for _, test := range tests {
		srv.Numbers = nil
		srv.Numbers = append(srv.Numbers, test.a, test.b, test.c)
		sum := srv.getNumbersSum()

		if sum != test.want {
			t.Errorf("Incorrect numbers sum result, got: %d, want: %d", sum, test.want)
		}
	}
}
