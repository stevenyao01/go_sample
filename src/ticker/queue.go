package main

/**
 * @Project: go_sample
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/9/8 下午12:15
 * @Description:
 */

// Queue 队列信息
type Queue struct{
	list *SingleList
}

// Init 队列初始化
func (q *Queue)Init()  {
	q.list = new(SingleList)
	q.list.Init()
}

// Size 获取队列长度
func (q *Queue) Size() uint{
	return q.list.Size
}

// Enqueue 进入队列
func (q *Queue) Enqueue(data interface{}) bool{
	return q.list.Append(&SingleNode{Data:data})
}

// Dequeue 出列
func (q *Queue) Dequeue() interface{}{
	node := q.list.Get(0)
	if node == nil{
		return nil
	}
	q.list.Delete(0)
	return node.Data
}

// Peek 查看队头信息
func (q *Queue)Peek() interface{}{
	node := q.list.Get(0)
	if node == nil{
		return nil
	}
	return node.Data
}

/**
 * @Description:
 * @Params:
 * @return:
 * @Date: 2020/9/17 上午10:01
 */
func QueueNew() (*Queue, error) {
	queue := Queue{}
	queue.Init()
	return &queue, nil
}
