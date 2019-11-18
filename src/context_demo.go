package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

}

//func NewContext(ctx context.Context) context.Context {
//	return context.WithValue(ctx, )
//}

// 可以使用Context让goroutine退出
// 可以避免goroutine内存泄漏
func WithCancelDemo() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// 当某项任务需要在规定时间内完成，如未完成则需要立即取消任务，并返回错误情况
func WithDeadlineDemo() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept!")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func WithTimeoutDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept!")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// 基于现有的context派生出一个新的context，新的带有函数指定的key与value
func WithValueDemo() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", k)
			return
		}
		fmt.Println("key", k, "not found.")
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Golang")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}
