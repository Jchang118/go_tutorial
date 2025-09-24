package data_structure_test

import (
    "go_tutorial/go_frame/data_structure"
    "fmt"
    "math/rand"
    "testing"
)

func TestQuickSort(t *testing.T) {
    for i := 0; i < 100; i++ { //测试100个case
        slice := make([]int, 20) //数组长度为20
        for j := 0; j < 20; j++ {
            slice[j] = rand.Intn(100)
        }
        data_structure.Partition(slice) //原地快速排序
        for j := 1; j < len(slice); j++ {
            if slice[j] < slice[j-1] { //没有从小到大的顺序排
                t.Fail()
            }
        } 
    }
}

func TestQuickSort2(t *testing.T) {
    arr := []int{4, 3, 6, 1, 27}
    data_structure.Partition(arr)
    fmt.Println(arr)
}

// go test ./data_structure -v -run=^TestQuickSort$ -count=1
// go test ./data_structure -v -run=^TestQuickSort2$ -count=1
