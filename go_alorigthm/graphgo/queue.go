package graphgo


type items []interface{}
type Queue struct{
	items items
}

func New(hint int)*Queue{
	return &Queue{
		items:make([]interface{},0,hint),
	}
}

func (queue *Queue)enqueue(ops ...interface{})error{
	queue.items = append(queue.items,ops...)
	return nil
}

func (queue *Queue)dequeue()interface{}{
	if len(queue.items) == 0{
		return nil
	}
	iEle := queue.items[0:1]
	queue.items = queue.items[1:]
	return iEle[0]
}

func (queue *Queue)isEmpty()bool{
	if len(queue.items)==0{
		return true
	}
	return false
}

func (queue *Queue) front()interface{}{
	if len(queue.items) == 0{
		return nil
	}
	return queue.items[0:1]
}

func (queue *Queue) Len()int{
	return len(queue.items)
}