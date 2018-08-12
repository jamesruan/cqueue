package cqueue

import (
	"sync"
	"sync/atomic"
)

type node struct {
	value interface{}
	next *node
}

type queue struct {
	head *node
	tail *node
	length uint32
	hlock *sync.Mutex
	tlock *sync.Mutex
}

type Queue interface {
	Length() uint32
	Enqueue(v interface{})
	Dequeue()(v interface{}, ok bool)
}

func New() Queue {
	node := new(node)
	return &queue{
		head: node,
		tail: node,
		hlock: new(sync.Mutex),
		tlock: new(sync.Mutex),
	}
}

func (q *queue) Length() uint32 {
	return atomic.LoadUint32(&q.length)
}

func (q *queue) Enqueue(v interface{}) {
	node := &node{
		value: v,
	}
	q.tlock.Lock()
	q.tail.next = node
	q.tail = node
	atomic.AddUint32(&q.length, 1)
	q.tlock.Unlock()
}

func (q *queue) Dequeue() (interface{}, bool) {
	q.hlock.Lock()
	node := q.head
	nh := node.next
	if nh == nil {
		// queue empty
		q.hlock.Unlock()
		return nil, false
	}
	v := nh.value
	q.head = nh
	atomic.AddUint32(&q.length, ^uint32(0))
	q.hlock.Unlock()
	return v, true
}
