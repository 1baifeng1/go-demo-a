package main

import (
	"fmt"
	"sort"
	"sync"
)

func say(text string) {
	fmt.Println(text)
}

func map_func() {
	// 创建map时，value可以是函数
	var funcMap = make(map[string]func(string))
	funcMap["say"] = say
	funcMap["say"]("hello world.")
}

func main() {
	map_func()
}

// 实现一个并发安全的KV缓存
type KV struct {
	count int
	keys  []string
	hash  map[string]interface{}
	lock  sync.RWMutex
}

// 新建KV缓存
func NewKV(preCapacity uint) *KV {
	return &KV{
		keys: make([]string, 0, int(preCapacity)),
		hash: make(map[string]interface{}, int(preCapacity)),
	}
}

// 添加KV键值对
func (this *KV) Set(k string, v interface{}) {
	this.lock.Lock()
	if _, ok := this.hash[k]; !ok {
		this.keys = append(this.keys, k)
		sort.Strings(this.keys)
		this.count++
	}
	this.hash[k] = v
	this.lock.Unlock()
}

// 获取数据长度
func (this *KV) Count() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.count
}

// 通过key检索value
func (this *KV) Get(k string) (interface{}, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	v, ok := this.hash[k]
	return v, ok
}

// 根据添加的顺序，返回有序的value切片
func (this *KV) Values() []interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	vals := make([]interface{}, this.count)
	for i := 0; i < this.count; i++ {
		vals[i] = this.hash[this.keys[i]]
	}
	return vals
}
