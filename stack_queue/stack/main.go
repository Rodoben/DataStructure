package main

import "stackpkg.com/stackpkg"

func main() {

	var S stackpkg.Stack

	S.Push(10)
	S.Push(20)
	S.DisplayStack()

	S.Pop()
	S.DisplayStack()
}
