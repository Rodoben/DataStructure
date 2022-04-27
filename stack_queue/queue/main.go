package main

import (
	"fmt"

	"queuePkg.com/queuePkg"
)

func main() {

	var Q queuePkg.Queue
	Q.Enqueue(10)
	Q.Enqueue(20)
	Q.Enqueue(30)
	//Q.Enqueue(20)

	//Q.DisplayQueue(Q)

	Q.Dequeue(Q)
	Q.Dequeue(Q)

	fmt.Println("***********************")
	Q.DisplayQueue()
}
