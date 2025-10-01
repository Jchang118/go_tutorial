package io

import (
	"io"
	"net"
	"os"
	"syscall"
)

// 常规的文件拷贝
func CopyFile(file string, fileSize int, w io.WriteCloser) {
	// 打开文件
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 读取文件内容
	data := make([]byte, fileSize)
	_, err = f.Read(data)
	if err != nil {
		panic(err)
	}

	// 把内容写入磁盘
	w.Write(data)
	w.Close() //确保内容全部Flush入磁盘
}

// 通过mmap拷贝文件
func CopyFileWithMmap(file string, fileSize int, w io.WriteCloser) {
	// 打开文件
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 通过mmap读取文件内容
	data, err := syscall.Mmap(int(f.Fd()), 0, fileSize, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	// 把内容写入磁盘
	w.Write(data)
	w.Close() //确保内容全部Flush入磁盘
}

// 把文件通过TCP连接发出去
func SendFile(file string, fileSize int, conn *net.TCPConn) {
	// 打开文件
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 读取文件内容
	data := make([]byte, fileSize)
	_, err = f.Read(data)
	if err != nil {
		panic(err)
	}

	// 把内容写给socket缓存
	conn.Write(data)
	conn.Close() //确保内容全部Flush给对方
}

// 通过系统调用Sendfile,把文件发送给网卡
func SendFileWithSyscall(file string, fileSize int, conn *net.TCPConn) {
	// 打开文件
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 通过系统调用Sendfile,把文件发送给网卡
	var offset int64 = 0
	outFile, _ := conn.File()
	_, err = syscall.Sendfile(int(outFile.Fd()), int(f.Fd()), &offset, fileSize)
	if err != nil {
		panic(err)
	}

	conn.Close()
}
