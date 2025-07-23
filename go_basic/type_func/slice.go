package main

import (
	"fmt"
	"strings"
)

func slice_init() {
	var s []int //切片声明, len=cap=0
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{} //初始化,len=cap=0
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3) //初始化,len=cap=3
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3, 5) //初始化,len=3,cap=5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{1, 2, 3, 4, 5} //初始化,len=cap=5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	fmt.Println("===========================")

	//二维切片初始的一种方式
	s2d := [][]int{
		{1},
		{2, 3}, //二维数组各行的列数是相等的,但二维切片各行的len可以不等
	}
	fmt.Printf("s2d len %d cap %d\n", len(s2d), cap(s2d))
	fmt.Printf("s2d[0] len %d cap %d\n", len(s2d[0]), cap(s2d[0]))
	fmt.Printf("s2d[1] len %d cap %d\n", len(s2d[1]), cap(s2d[1]))
	fmt.Println("===========================")
}

func slice_append() {
	arr := make([]int, 3, 6)
	brr := append(arr, 8) //arr和brr共享底层数组,但它们的len不同
	brr[0] = 9

	fmt.Printf("arr[0]=%d, cap of arr %d, len of arr %d\n", arr[0], cap(arr), len(arr))
	fmt.Printf("brr[0]=%d, cap of brr %d, len of brr %d\n", brr[0], cap(brr), len(brr))
	
	s := make([]int, 3, 5)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	} // s=[1,2,3]
	fmt.Printf("s[0] address %p, s=%v\n", &s[0], s)
	/*
		capacity还够用,直接把追加的元素放到预留的内存空间上
	*/
	s = append(s, 4, 5) //可以一次append多个元素
	fmt.Printf("s[0] address %p, s=%v\n", &s[0], s)
	/*
		capacity不够用了,得申请一片新的内存,把老数据先拷贝过来,在新内存上执行append操作
	*/
	s = append(s, 6)
	fmt.Printf("s[0] address %p, s=%v\n", &s[0], s)
	fmt.Println("===========================")
}
	
func sub_slice() {
	/*
		截取一部分,创造子切片,此时子切片与母切片(或母数组)共享底层内存空间,母切片的capacity子切片可能直接用
	*/
	s := make([]int, 3, 5)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	} //s=[1,2,3]
	fmt.Printf("s[1] address %p\n", &s[1])
	sub_slice := s[1:3] //从切片创造子切片, len=cap=2
	fmt.Printf("len %d cap %d\n", len(sub_slice), cap(sub_slice))
	/*
		母切片的capacity还允许子切片执行append操作
	*/
	sub_slice = append(sub_slice, 6, 7) //可以一次append多个元素
	sub_slice[0] = 8
	fmt.Printf("s=%v, sub_slice=%v, s[1] address %p. sub_slice[0] address %p\n", s, sub_slice, &s[1], &sub_slice[0])
	/*
		母切片的capacity用完了,子切片再执行append就得申请一片新的内存,把老数据先拷贝过来,在新内存上执行append操作.此时的append操作跟母切片没有任何关系
	*/
	sub_slice = append(sub_slice, 8)
	sub_slice[0] = 9
	fmt.Printf("s=%v, sub_slice=%v, s[1] address %p. sub_slice[0] address %p\n", s, sub_slice, &s[1], &sub_slice[0])

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr[1] address %p\n", &arr[1])
	sub_slice = arr[1:3] //从数组创造子切片,len=2. cap=4
	fmt.Printf("len %d cap %d\n", len(sub_slice), cap(sub_slice))
	/*
		母数组的capacity还允许子切片执行append操作
	*/
	sub_slice = append(sub_slice, 6, 7) //可以一次append多个元素
	sub_slice[0] = 8
	fmt.Printf("arr=%v, sub_slice=%v, arr[1] address %p, sub_slice[0] address %p\n", arr, sub_slice, &arr[1], &sub_slice[0])
	/*
		母数组的capacity用完了,子切片再执行append就得申请一片新的内存,把老数据先拷贝过来,在新内存上执行append操作.此时的append操作跟母数组没有任何关系
	*/
	sub_slice = append(sub_slice, 8)
	sub_slice[0] = 9
	fmt.Printf("arr=%v, sub_slice=%v, arr[1] address %p, sub_slice[0] address %p\n", arr, sub_slice, &arr[1], &sub_slice[0])
	
	array := [...]int{1, 2, 3, 4, 5}
	brr := array[:]		//截取数组,得到的是切片
	crr := brr[1:2:4]
	fmt.Printf("len(crr)=%d, cap(crr)=%d\n", len(crr), cap(crr))

	fmt.Println("===========================")
}

// 清空slice
func clear_slice(arr *[]int) {
	*arr = []int{}
}

// 探究capacity扩容规律
func expansion() {
	s := make([]int, 0, 3)
	prevCap := cap(s)
	for i := 0; i < 2000; i++ {
		s = append(s, i)
		currCap := cap(s)
		if currCap > prevCap {
			//每次扩容都是扩到原先的2倍
			fmt.Printf("capacity从%d变到%d\n", prevCap, currCap)
			prevCap = currCap
		}
	}
	fmt.Println("===========================")
}

// go语言函数传参,传的都是值,即传切片会把切片的{arrayPointer, len, cap}这3个属性拷贝一份传进来.
// 由于传的是底层数组的指针,所以可以直接修改底层数组里的元素

