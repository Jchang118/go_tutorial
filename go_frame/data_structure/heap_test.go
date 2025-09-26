package data_structure_test

import (
    "container/heap"
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

func TestStdHeap(t *testing.T) {
    pq := make(data_structure.PriorityQueue[int], 0, 10)
    pq.Push(&data_structure.Item[int]{Info: "A", Value: 3}) //往数组里添加元素
    pq.Push(&data_structure.Item[int]{Info: "B", Value: 2}) 
    pq.Push(&data_structure.Item[int]{Info: "C", Value: 4}) 
    heap.Init(&pq)                                                  //根据数组中的元素构建堆
    heap.Push(&pq, &data_structure.Item[int]{Info: "D", Value: 6})  //通过heap添加元素
    //通过不断删除堆顶,来实现堆排序
    for pq.Len() > 0 {
        fmt.Println(heap.Pop(&pq)) //删除堆顶元素(即最小的元素)
    }
}

// go test ./data_structure -v -run=^TestHeap$ -count=1
// go test ./data_structure -v -run=^TestStdHeap$ -count=1
