package stackpkg

import "fmt"

type Stack []int

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
func (s *Stack) DisplayStack() {
	for i, v := range *s {
		fmt.Println(i, v)
	}
}


func (s *Stack)Delelte{
	
}