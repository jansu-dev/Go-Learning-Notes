package computer

import "fmt"

type computer struct{}

func (com computer) work() {
	fmt.Println("computer is working!")
}

func (com computer) finish() {
	fmt.Println("computer is finish!")
}
