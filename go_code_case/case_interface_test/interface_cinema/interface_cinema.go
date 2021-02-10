package ciname

import "fmt"

type cinema struct{}

func (ca cinema) work() {
	fmt.Println("cinema is working!")
}

func (ca cinema) finish() {
	fmt.Println("cinema is finish!")
}
