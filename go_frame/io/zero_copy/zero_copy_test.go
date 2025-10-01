package io_test

import (
	"fmt"
	dqq_io "go_tutorial/go_frame/io/zero_copy"
	"io"
	"net"
	"os"
	"testing"
	"time"
)

var (
	file = "/Users/liangchengchang/2025/goprojects.zip"
)

func getFileSize(file string) int64 {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return stat.Size()
}

func getWriter(outFile string) io.WriteCloser {
	w, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return w
}

func getTcpConn(host string) *net.TCPConn {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", host)
	conn, _ := net.DialTCP("tcp4", nil, tcpAddr)
	return conn
}

func TestCopyFile(t *testing.T) {
	fileSize := getFileSize(file)
	w := getWriter("/Users/liangchengchang/2025/copy.zip")
	begin := time.Now()
	dqq_io.CopyFile(file, int(fileSize), w)
	fmt.Printf("copy file time spent %d ms\n", time.Since(begin).Milliseconds())
}

func TestCopyFileWithMmap(t *testing.T) {
	fileSize := getFileSize(file)
	w := getWriter("/Users/liangchengchang/2025/mmap_copy.zip")
	begin := time.Now()
	dqq_io.CopyFileWithMmap(file, int(fileSize), w)
	fmt.Printf("copy file with mmap time spent %d ms\n", time.Since(begin).Milliseconds())
}

func TestSendFile(t *testing.T) {
	fileSize := getFileSize(file)
	conn := getTcpConn("127.0.0.1:5678")
	begin := time.Now()
	dqq_io.SendFile(file, int(fileSize), conn)
	fmt.Printf("send file time spent %d ms\n", time.Since(begin).Milliseconds())
}

func TestSendFileWithSyscall(t *testing.T) {
	fileSize := getFileSize(file)
	conn := getTcpConn("127.0.0.1:5678")
	begin := time.Now()
	dqq_io.SendFileWithSyscall(file, int(fileSize), conn)
	fmt.Printf("send file with syscall time spent %d ms\n", time.Since(begin).Milliseconds())
}

// go test -v ./io/zero_copy -run=^TestCopyFile$ -count=1
// go test -v ./io/zero_copy -run=^TestCopyFileWithMmap$ -count=1
// go test -v ./io/zero_copy -run=^TestSendFile$ -count=1
// go test -v ./io/zero_copy -run=^TestSendFileWithSyscall$ -count=1
