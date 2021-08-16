package queue

type Queue struct {
	max        int
	front      int
	rear       int
	noElements int
	array      []rune
}

func NewQueue(size int) *Queue {
	return &Queue{
		max:   size,
		front: 0,
		rear:  -1,
		array: make([]rune, size),
	}
}

func (q *Queue) Insert(c rune) {
	if q.noElements == q.max {
		return
	}
	if q.rear == q.max-1 {
		q.rear = -1
	}
	q.rear++
	q.array[q.rear] = c
	q.noElements++
}

func (q *Queue) Remove() rune {
	if q.noElements == 0 {
		return '0'
	}
	if q.front == q.max {
		q.front = 0
		q.noElements--
		return q.array[q.front]
	}
	q.noElements--
	q.front++
	return q.array[q.front]
}

func (q *Queue) PeekFront() rune {
	return q.array[q.front]
}
