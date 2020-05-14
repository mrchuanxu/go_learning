package queue_test

// queue 先入先出
// 循环队列实现

type queue struct {
	Tail int64
	Head int64
	Size int64
	Q    []string
}

// 初始化队列
func NewQ() *queue {
	q := queue{
		Tail: 0,
		Head: 0,
		Size: 100,
	}
	q.Q = make([]string, q.Size)
	return &q
}

func (q *queue) InQueue(val string) bool {
	if (q.Tail+1)%q.Size == q.Head {
		return false
	}
	q.Q[q.Tail] = val
	q.Tail = (q.Tail + 1) % q.Size
	return true
}

func (q *queue) DeQueue() bool {
	if q.Tail == q.Head {
		return false
	}
	q.Head = (q.Head + 1) % q.Size
	return true
}

func (q *queue) Front() string {
	if q.Tail == q.Head {
		return ""
	}
	return q.Q[q.Head]
}
