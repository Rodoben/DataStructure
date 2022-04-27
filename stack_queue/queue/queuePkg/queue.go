package queuePkg

import "fmt"

type Queue []int

func (q *Queue) Enqueue(element int) {
	*q = append(*q, element)
	//eq = append(eq, element)
	fmt.Println("Enqueued:", element)
	fmt.Println("EQ:", *q)

}

func (q *Queue) Dequeue() []int {
	index := (*q)[0]
	fmt.Println("Dequeued:", index)

	a := (*q)[1:]
	fmt.Println(a)
	return a

}

func (q *Queue) DisplayQueue() {

	fmt.Println("ds:", *q)
	for i, v := range *q {
		fmt.Println(i, v)

	}
}
