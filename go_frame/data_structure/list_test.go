package data_structure_test

import (
    "container/list"
    "go_tutorial/go_frame/data_structure"
    "fmt"
    "testing"
)

func TestList(t *testing.T) {
    lst := new(data_structure.DoubleList[int])
    lst.PushBack(1)     // 1
    lst.PushBack(2)     // 1 -> 2
    lst.PushFront(3)    // 3 -> 1 -> 2
    lst.PushFront(4)    // 4 -> 3 -> 1 -> 2

    third := lst.Get(2)         //第三个元素是1
    lst.InsertAfter(8, third)  // 4 -> 3 -> 1 -> 8 -> 2
    lst.InsertBefore(9, third) // 4 -> 3 -> 9 -> 1 -> 8 -> 2

    fmt.Printf("链表中共有%d个元素\n", lst.Length)
    lst.Traverse()
    lst.ReverseTraverse()
}

func TestStdList(t *testing.T) {
    lst := list.New()
    lst.PushBack(1)     // 1
    lst.PushBack(2)     // 1 -> 2
    lst.PushFront(3)    // 3 -> 1 -> 2
    lst.PushFront(4)    // 4 -> 3 -> 1 -> 2

    third := lst.Front().Next().Next()  //第三个元素是1
    lst.InsertAfter(8, third)           // 4 -> 3 -> 1 -> 8 -> 2
    lst.InsertBefore(9, third)          // 4 -> 3 -> 9 -> 1 -> 8 -> 2
    lst.Remove(third)                   // 4 -> 3 -> 9 -> 8 -> 2

    fmt.Printf("链表中共有%d个元素\n", lst.Len())
    data_structure.TraverseList(lst)
    data_structure.ReverseList(lst)
}

// go test ./data_structure -v -run=^TestList$ -count=1
// go test ./data_structure -v -run=^TestStdList$ -count=1
