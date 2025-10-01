package io

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/exp/mmap"
)

// 随机选择位置读文件
func ReadDiskRandomly(file string) {
	t0 := time.Now()
	defer func() {
		fmt.Printf("随机读文件 %dms\n", time.Since(t0).Milliseconds())
	}()

	// 打开文件
	fin, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1)
	for i := 0; i < len(arr); i++ {
		offset := int64(index[i])
		fin.ReadAt(buffer, offset) // 随机读取文件某个位置上的字节
	}
}

// 随机选择位置读文件
func GolangMmap(file string) {
	t0 := time.Now()
	defer func() {
		fmt.Printf("GolangMmap 随机读文件 %dms\n", time.Since(t0).Milliseconds())
	}()

	// 打开文件
	ra, err := mmap.Open(file)
	if err != nil {
		panic(err)
	}
	defer ra.Close()

	buffer := make([]byte, 1)
	for i := 0; i < len(arr); i++ {
		offsest := int64(index[i])
		ra.ReadAt(buffer, offsest) // 随机读取文件某个位置上的字节
	}
}

// 把文件先加载到内存再随机读
func LoadAndReadRandomly(file string) {
	t0 := time.Now()
	defer func() {
		fmt.Printf("把文件先加载到内存再随机读 %dms\n", time.Since(t0).Milliseconds())
	}()

	// 打开文件
	fin, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	// 把文件加载到内存
	if n, err := fin.Read(arr); err != nil || n != ARRAY_SIZE {
		log.Fatal(n, err)
	}

	// 随机读内存
	for _, i := range index {
		_ = arr[i]
	}
}

/*
随机读写mmap虽然没有随机读取内存快,但mmap的优势在于:
1. 基于磁盘的读写单位是block(一般大小为4KB),而基于内存的读写单位是地址(虽然内存的管理与分配单位是4KB).
2. mmap向应用程序提供的内存访问接口是内存地址连续的,但是对应的磁盘文件的block可以不是地址连续的.
3. mmap提供的内存空间是虚拟空间(虚拟内存),而不是物理空间(物理内存),因此完全可以分配远远大于物理内存大小的虚拟空间(例如16G内存主机分配1000G的mmap内存空间).
*/
