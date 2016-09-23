package bowling

import (
	"fmt"
	"testing"
)

func scoreChecker(input []Frame, expectedScore int, expectedError error) error {
	score, err := GetScore(input)

	if err != expectedError && !(err != nil && expectedError != nil && err.Error() == expectedError.Error()) {
		return fmt.Errorf("Score error : %+v, expected %+v", err, expectedError)
	}
	if score != expectedScore {
		return fmt.Errorf("Score : %d, expected %d", score, expectedScore)
	}
	return nil
}

func TestScore(t *testing.T) {
	input := []Frame{{5, 4}, {5, 4}, {5, 4}, {5, 4}, {5, 4}, {5, 4}, {5, 4}, {5, 4}, {5, 4}, {5, 4}}
	expected := 90
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestNegScore(t *testing.T) {
	input := []Frame{{1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, -1}}
	expected := 0
	errExpected := fmt.Errorf("Lesser than 0 found")
	if err := scoreChecker(input, expected, errExpected); err != nil {
		t.Fatalf("%+v\n", err)
	}
}
func TestSize(t *testing.T) {
	input := []Frame{{1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	expected := 0
	errExpected := fmt.Errorf("Less than 10 frame found")
	if err := scoreChecker(input, expected, errExpected); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSumFrame(t *testing.T) {
	input := []Frame{{10, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	expected := 0
	errExpected := fmt.Errorf("Score more than 10 on single frame")
	if err := scoreChecker(input, expected, errExpected); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSpare(t *testing.T) {
	input := []Frame{{5, 5}, {4, 4}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 10 + 4 + 4 + 4
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestStrike(t *testing.T) {
	input := []Frame{{10, 0}, {4, 4}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	expected := 10 + 4 + 4 + 4 + 4
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestStrikeLast(t *testing.T) {
	input := []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {10, 0}}
	expected := 10
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func TestSpareLast(t *testing.T) {
	input := []Frame{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {5, 5}}
	expected := 10
	if err := scoreChecker(input, expected, nil); err != nil {
		t.Fatalf("%+v\n", err)
	}
}
