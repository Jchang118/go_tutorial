package data_structure

import "container/ring"

// 基于Ring的滑动窗口
type SlideWindow struct {
    n       int
    ring    *ring.Ring
}

func NewSlideWindow(n int) *SlideWindow {
    return &SlideWindow{
        n:      n,
        ring:   ring.New(n),
    }
}

// 添加元素
func (w *SlideWindow) Push(data float64) {
    w.ring.Value = data     //给当前元素赋值,沿未赋值时Value为nil
    w.ring = w.ring.Next()  //游标指向下一个元素
}

// 统计滑动窗口内所有元素的均值
func (w *SlideWindow) Mean() float64 {
    var sum, count float64
    // 遍历Ring,通过Do()函数指定如何处理单个元素
    w.ring.Do(func(a any) {
        if a != nil { //不统计没有赋值的元素
            count += 1
            sum += a.(float64)
        }
    })
    return sum / count
}
