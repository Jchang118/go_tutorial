package data_structure_test

import (
    "go_tutorial/go_frame/data_structure"
    "fmt"
    "testing"
    "slices"
)

func TestJaccardTimeConsuming(t *testing.T) {
    l1 := []string{"go", "分布式", "mysql", "搞笑", "并发编程", "服务器"}
    l2 := []string{"c#", "AI", "mysql", "篮球", "并发编程", "服务器"}
    fmt.Println(data_structure.JaccardTimeConsuming(l1, l2))
}

func TestJaccardForSorted(t *testing.T) {
    l1 := []string{"go", "分布式", "mysql", "搞笑", "并发编程", "服务器"}
    l2 := []string{"c#", "AI", "mysql", "篮球", "并发编程", "服务器"}
    slices.Sort(l1)
    slices.Sort(l2)
    fmt.Println(data_structure.JaccardForSorted(l1, l2))
}

// go test ./data_structure -v -run=^TestJaccardTimeConsuming$ -count=1
// go test ./data_structure -v -run=^TestJaccardForSorted$ -count=1
