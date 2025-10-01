package io_test

import (
	"go_tutorial/go_frame/io"
	"testing"
)

func TestRandomIO(t *testing.T) {
	file := "../data/arr.bin"
	io.InitArray()
	io.ReadDiskRandomly(file)
	io.GolangMmap(file)
	io.LoadAndReadRandomly(file)
}

// go test -v ./io -run=^TestRandomIO$ -count=1
