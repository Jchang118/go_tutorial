package data_structure_test

import (
    "go_tutorial/go_frame/data_structure"
    "math/rand"
    "slices"
    "sort"
    "testing"
)

func TestBinarySearch(t *testing.T) {
    const L = 100   //数组的长度L
LOOP:
    for i := 0; i < 100; i++ { //测试100个case
        arr := make([]int, L)
        for j := 0; j < L; j++ {
            arr[j] = rand.Intn(L)
        }
        sort.Slice(arr, func(i, j int) bool {return arr[i] < arr[j]})

        target := rand.Intn(L)
        index := data_structure.BinarySearch[int](arr, target)
        if index >= 0 {
            if arr[index] != target {
                t.Fail()
                break LOOP
            }
        } else {
            for _, ele := range arr {
                if ele == target {
                    t.Fail()
                    break LOOP
                }
            }
        }
    }
}

func TestBinarySearch1(t *testing.T) {
    arr := []int{1, 3, 6, 8, 10}
    targets := []int{1, 6, 10}
    for _, target := range targets {
        index := data_structure.BinarySearch[int](arr, target)
        if index < 0 {
            t.Fatalf("case %d 测试失败\n", target)
        }
    }
}

// B+树里每个节点有1200个子节点,在1200个元素时查找特定元素需要多长时间?
func BenchmarkBinarySearchBPlusTree(b *testing.B) {
    const M = 1200
    array := make([]int, M)
    for i := 0; i < M; i++ {
        array[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ { //55ns
        target := rand.Intn(M)
        data_structure.BinarySearch(array, target)
    }
}

// 测试二分区间查找
func TestBinarySearch4Section(t *testing.T) {
    const L = 100
    for c := 0; c < 30; c++ { //进行多轮严格的测试　
        arr := make([]float64, 0, L)
        for i := 0; i < L+c; i++ { //每轮的切片长度也有所变化
            arr = append(arr, rand.Float64())
        }
        slices.Sort(arr) //排序
        var target float64

        //先测试2个越界的情况
        target = arr[0] - 1.0
        if data_structure.BinarySearch4Section(arr, target) != 0 {
            t.Fail()
        }
        target = arr[len(arr)-1] + 1.0
        if data_structure.BinarySearch4Section(arr, target) != len(arr) {
            t.Fail()
        }
         
        // 每个分割点,及2个分割点中间的值都测一下
        target = arr[0]
        if data_structure.BinarySearch4Section(arr, target) != 0 {
            t.Fail()
        }

        for i := 0; i < len(arr)-1; i++ {
            target = (arr[i] + arr[i+1]) / 2
            if data_structure.BinarySearch4Section(arr, target) != i+1 {
                t.Fail()
            }
            target = arr[i+1]
            if data_structure.BinarySearch4Section(arr, target) != i+1 {
                t.Fail()
            }
        }
    }
}

// go test ./data_structure -v -run=^TestBinarySearch$ -count=1
// go test ./data_structure -v -run=^TestBinarySearch1$ -count=1
// go test ./data_structure -bench=^BenchmarkBinarySearchBPlusTree$ -run=^$ -count=1
// go test ./data_structure -v -run=^TestBinarySearch4Section$ -count=1
