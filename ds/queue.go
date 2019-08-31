package ds

import (
	"container/list"
	"reflect"
	"sync"
)

type Queue struct {
	data *list.List

	enableLock bool
	mu         *sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		data: list.New(),
		mu:   &sync.Mutex{},
	}
}

func (queue *Queue) EnableLock(b bool) *Queue {
	queue.enableLock = b
	return queue
}

func (queue *Queue) Push(v interface{}) {
	if queue.enableLock {
		queue.mu.Lock()
		defer queue.mu.Unlock()
	}
	queue.data.PushBack(v)
}

func (queue *Queue) PushN(v interface{}) {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice || vv.Len() == 0 {
		return
	}
	if queue.enableLock {
		queue.mu.Lock()
		defer queue.mu.Unlock()
	}
	for i := 0; i < vv.Len(); i++ {
		queue.data.PushBack(vv.Index(i).Interface())
	}
}

func (queue *Queue) Pop() interface{} {
	if queue.enableLock {
		queue.mu.Lock()
		defer queue.mu.Unlock()
	}
	if queue.data.Len() == 0 {
		return nil
	}
	return queue.data.Remove(queue.data.Front())
}

func (queue *Queue) PopN(n int) []interface{} {
	if queue.enableLock {
		queue.mu.Lock()
		defer queue.mu.Unlock()
	}
	if queue.data.Len() < n {
		n = queue.data.Len()
	}
	if n <= 0 {
		return nil
	}
	vs := make([]interface{}, n)
	for i := 0; i < n; i++ {
		vs[i] = queue.data.Remove(queue.data.Front())
	}
	return vs
}

func (queue *Queue) Len() int {
	if queue.enableLock {
		queue.mu.Lock()
		defer queue.mu.Unlock()
	}
	return queue.data.Len()
}
