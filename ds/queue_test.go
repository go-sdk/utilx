package ds

import (
	"math/rand"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	queue.EnableLock(true)
	t.Logf("len0: %d", queue.Len())
	for i := 1; i <= 5; i++ {
		t.Logf("len1: %d", queue.Len())
		t.Logf("push: %d", i)
		queue.Push(i)
		array := make([]int, rand.Intn(10))
		for j := 0; j < len(array); j++ {
			array[j] = rand.Intn(10)
		}
		t.Logf("push %d: %v", len(array), array)
		queue.PushN(array)
		t.Logf("len2: %d", queue.Len())
		if i%2 == 0 {
			t.Logf("pop: %v", queue.Pop())
		} else {
			t.Logf("pop %d: %v", i, queue.PopN(i))
		}
		t.Logf("len3: %d", queue.Len())
	}
	t.Logf("len0: %d", queue.Len())
	t.Logf("pop %d: %v", queue.Len(), queue.PopN(queue.Len()))
	t.Logf("len0: %d", queue.Len())

	queue.PushN(1)
	t.Logf("pop: %v", queue.Pop())
	t.Logf("pop: %v", queue.PopN(0))
	t.Logf("pop: %v", queue.PopN(1))
}
