package data_structure

import (
    "cmp"
)

// 二分查找.返回target在arr中的下标,如果不存在则返回-1,如果target中存在多个只返回其中某一个的下标.这里使用泛型可支持多种数据类型
func BinarySearch[T cmp.Ordered](arr []T, target T) int {
    begin := 0
    end := len(arr) - 1
    for begin <= end { //之所以是<=,而不是<,是因为区间内只剩下一个元素时也应该跟target进行比较,而不是直接返回-1
        middle := (begin + end) / 2
        if arr[middle] == target {
            return middle
        } 
        if arr[middle] < target {
            begin = middle + 1
        } else {
            end = middle - 1
        }
    }
    return -1
}

// 二分法查找数组中>=target的最小的元素下标.arr是单调递增的(里面不能存在重复的元素),如果target比arr的最后一个元素还大,则返回最后arr的长度;如果target比arr的第一个元素还小,则返回0
func BinarySearch4Section[T cmp.Ordered](arr []T, target T) int {
    if len(arr) == 0 {
        return -1
    }
    begin, end := 0, len(arr)-1

    for {
        //arr[begin]在target后面
        if arr[begin] >= target {
            return begin
        }
        //arr[end]在target前面
        if arr[end] < target {
            return end + 1
        }

        //二分查找法
        middle := (begin + end) / 2
        if arr[middle] > target {
            end = middle - 1 //arr[end]可能会跑到target前面
        } else if arr[middle] < target {
            begin = middle + 1 //arr[begin]可能会跑到target后面
        } else {
            return middle
        }
    }
}
