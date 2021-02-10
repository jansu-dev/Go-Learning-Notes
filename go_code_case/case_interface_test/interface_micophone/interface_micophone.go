package micophone

import "fmt"

type micophone struct{}

func (com micophone) work() {
	fmt.Println("micophone is working!")
}

func (com micophone) finish() {
	fmt.Println("micophone is finish!")
}
