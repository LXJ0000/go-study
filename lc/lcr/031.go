package main

import (
	"fmt"
)

type LRUCache struct {
	capacity   int
	head, tail *LRUNode         // 双向链表
	cache      map[int]*LRUNode // 字典 key - node
}

type LRUNode struct {
	key, val   int
	prev, next *LRUNode
}

func NewLRUNode(key, val int) *LRUNode {
	return &LRUNode{
		key: key,
		val: val,
	}
}

func Constructor1(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		head:     NewLRUNode(0, 0),
		tail:     NewLRUNode(0, 0),
		cache:    make(map[int]*LRUNode),
	}
	// 初始化双向链表
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		// 如果存在，则返回val,更新缓存：放置队头
		// 1. 删除节点
		// 2. 头插
		lru.RemoveNode(node)
		lru.AddToHead(node)
		return node.val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	node, ok := lru.cache[key]
	if ok {
		// key 存在，则修改val，更新缓存：放置队头
		node.val = value
		lru.RemoveNode(node)
		lru.AddToHead(node)
		return
	}
	// 不存在，构造节点，判断缓存大小，如果超过则删除队尾节点，再将节点头插
	node = &LRUNode{key: key, val: value}
	if len(lru.cache) == lru.capacity {
		lru.RemoveNode(lru.tail.prev)
	}
	lru.AddToHead(node)
}

func (lru *LRUCache) AddToHead(node *LRUNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node

	lru.cache[node.key] = node
}

func (lru *LRUCache) RemoveNode(node *LRUNode) {
	delete(lru.cache, node.key)
	node.next.prev = node.prev
	node.prev.next = node.next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor1(5)
	obj.Put(1, 2)

	param1 := obj.Get(1)
	fmt.Println(param1)
}
