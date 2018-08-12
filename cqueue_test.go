package cqueue

import (
	"testing"
)

func BenchmarkEnqueueAndDequeue(b *testing.B) {
	q := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Enqueue(0)
		}
	})
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Dequeue()
		}
	})
}

