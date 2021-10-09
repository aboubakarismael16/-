package main

import "fmt"

type Queue []string

func (q *Queue) Enqueue(str string)  {
	*q = append(*q, str)
}

func (q *Queue) Dequeue() (string, bool)  {
	if q.IsEmpty() {
		return "", false
	} else {
		elemen := (*q)[0]
		*q = (*q)[1:]
		return elemen, true
	}
}

func (q *Queue) IsEmpty() bool  {
	return len(*q) == 0
}

func main()  {
	var q Queue

	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue("3")
	q.Enqueue("4")

	fmt.Println(q)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	fmt.Println(q)

}