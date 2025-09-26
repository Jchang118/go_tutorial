package data_structure_test

import (
    //"container/heap"
    "go_tutorial/go_frame/data_structure"
    "fmt"
    "testing"
)

func TestHeap(t *testing.T) {
    h := data_structure.NewHeap[int]([]int{50, 20, 49, 15, 30, 62})
    h.Build()
    h.Push(5)
    //堆排序
    for h.Size() > 0 {
        top, _ := h.Pop()
        fmt.Println(top)
    }
}

// go test ./data_structure -v -run=^TestHeap$ -count=1
