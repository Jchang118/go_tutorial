package projectprepare_test

import (
	projectprepare "go_tutorial/go_basic/project_prepare"
	"testing"
)

func TestCopySlice(t *testing.T) {

	var src, dest []int16
	src = []int16{1, 2, 3, 4}
	var c, n int

	c = len(src) - 1
	dest = make([]int16, c)
	n = projectprepare.CopySlice(dest, src)
	if n != c {
		t.Errorf("c=%d, n %d", c, n)
	}
	for i := 0; i < n; i++ {
		if dest[i] != src[i] {
			t.Errorf("c=%d, i=%d, dest %d src %d", c, i, dest[i], src[i])
		}
	}

	c = len(src)
	dest = make([]int16, c)
	n = projectprepare.CopySlice(dest, src)
	if n != len(src) {
		t.Errorf("c=%d, n %d", c, n)
	}
	for i := 0; i < n; i++ {
		if dest[i] != src[i] {
			t.Errorf("c=%d, i=%d, dest %d src %d", c, i, dest[i], src[i])
		}
	}

	c = len(src) + 1
	dest = make([]int16, c)
	n = projectprepare.CopySlice(dest, src)
	if n != len(src) {
		t.Errorf("c=%d, n %d", c, n)
	}
	// dest[0]++
	// dest[1]++
	for i := 0; i < n; i++ {
		if dest[i] != src[i] {
			t.Fatalf("c=%d, i=%d, dest %d src %d", c, i, dest[i], src[i])
		}
	}
}

func BenchmarkCopySlice(b *testing.B) {
	src := make([]int8, 10000)
	dest := make([]int8, 10000)
	b.ResetTimer() //开始计时
	for i := 0; i < b.N; i++ {
		projectprepare.CopySlice(dest, src)
	}
}

func BenchmarkStdCopySlice(b *testing.B) {
	src := make([]int8, 10000)
	dest := make([]int8, 10000)
	b.ResetTimer() //开始计时
	for i := 0; i < b.N; i++ {
		copy(dest, src)
	}
}

// go test -v ./project_prepare -run=TestCopySlice$ -count=1
// got test ./project_prepare -run=^$ -bench=CopySlice$ -count=1
// go test -cover $dir 只能给出$dir目录的整体单测覆盖率
// go test ./project_prepare -coverprofile=data/test_cover
// go test ./project_prepare -coverprofile=data/test_cover -covermode=count
// covermode可以取3个值: set 每个语句是否执行- 默认值, count 每个语句执行了几次,鼠标悬停在语句上显示执行的次数, atomic 类似于count,但表示的是并行程序中的精确计数 
// go tool cover -func=data/test_cover 输出每一个函数的覆盖率
// go tool cover -html=data/test_cover 细化到每一行代码的覆盖情况
