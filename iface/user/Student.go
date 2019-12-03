package user

import "fmt"

type Student struct {
}

func New() Student {
	return Student{}
}
func (s *Student) Say() {
	fmt.Println("hello world!")
}
