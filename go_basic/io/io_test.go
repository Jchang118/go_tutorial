package io_test

import (
	"go_tutorial/go_basic/io"
	"testing"
	"time"
	"fmt"
)

func TestWriteFile(t *testing.T) {
	io.WriteFile()
}

func TestReadFile(t *testing.T) {
	io.ReadFile()
}

func TestWriteFileWithBuffer(t *testing.T) {
	io.WriteFileWithBuffer()
}

func TestReadFileWithBuffer(t *testing.T) {
	io.ReadFileWithBuffer()
}

func TestBufferedFileWriter(t *testing.T) {
	t1 := time.Now()
	io.WriteDirect("../data/no_buffer.txt")
	t2 := time.Now()
	io.WriteWithBuffer("../data/with_buffer.txt")
	t3 := time.Now()
	fmt.Printf("不用缓冲耗时%dms,用缓冲耗时%dms\n", t2.Sub(t1).Milliseconds(), t3.Sub(t2).Milliseconds())
}

func TestCreateFile(t *testing.T) {
	io.CreateFile("../data/poem.txt")
}

func TestWalkDir(t *testing.T) {
	io.WalkDir("../data")
}

func TestSplitFile(t *testing.T) {
	imgFile := "../img/大乔乔好课.png"
	io.SplitFile(imgFile, "../img/图像分割", 4)
}

func TestMergeFile(t *testing.T) {
	io.MergeFile("../img/图像分割", "../img/图像合并.png")
}

func TestCompress(t *testing.T) {
	io.Compress("../img/大乔乔好课.png", "../img/大乔乔好课.png.gzip", io.GZIP)
	io.Decompress("../img/大乔乔好课.png.gzip", "../data/大乔乔好课.png", io.GZIP)
}

func TestJson(t *testing.T) {
	io.JsonSerialize()
}
