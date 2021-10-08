package a03frogjmp

import "testing"

var testData = []struct {
    x   int
    y   int
    d   int
    out int
}{
    {10, 85, 30, 3},
    {10, 234_342_345, 8, 29_292_792},
    {10, 234_342_345, 1, 234_342_335},
    {10, 2_234_342_345, 1, 2_234_342_335},
    {10, 2_234_342_345, 1_000_000_000, 3},
    {100, 100, 2, 0},
}

func TestFrogJmp(t *testing.T) {
    for _, td := range testData {
        if td.out != SolutionFrogJmp(td.x, td.y, td.d) {
            t.Errorf("FAIL!")
        }
    }
}
