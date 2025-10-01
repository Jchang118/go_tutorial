package io_test

import (
	"go_tutorial/go_frame/io"
	"testing"
)

func TestSequentialIO(t *testing.T) {
	io.InitArray()
	io.ReadRamSequentially()
	io.ReadRamRandomly()
	io.WriteRamSequentially()
	io.WriteRamRandomly()
	io.WriteDiskSequentially()
	io.ReadDiskSequentially()
	// os.Remove("../data/arr.bin")
}

// go test -v ./io -run=^TestSequentialIO$ -count=1
