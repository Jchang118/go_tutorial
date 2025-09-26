package data_structure_test

import (
    "go_tutorial/go_frame/data_structure"
    "fmt"
    "testing"
)

func TestRing(t *testing.T) {
    window := data_structure.NewSlideWindow(5)
    for i := 0; i < 10; i++ {
        window.Push(float64(i))
        fmt.Printf("%d %f\n", i, window.Mean())
    }
}

// go test ./data_structure -v -run=^TestRing$ -count=1
