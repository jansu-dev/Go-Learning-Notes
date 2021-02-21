package main

import "fmt"

type cinema struct{}

// the work function of cinema
func (ca cinema) work() {
	fmt.Println("cinema is working!")
}

// the finish function of cinema
func (ca cinema) finish() {
	fmt.Println("cinema is finish!")
}

type micophone struct{}

// the work function of micophone
func (com micophone) work() {
	fmt.Println("micophone is working!")
}

// the finish function of micophone
func (com micophone) finish() {
	fmt.Println("micophone is finish!")
}

// the computer struct which is belong to main function
type computer struct{}

// that usb is the interface ,which awllow implemented interface of others to be used, work for computer
type usb interface {
	work()
	finish()
}

func (com computer) start(usb usb) {
	usb.work()
	usb.finish()
}

func main() {
	computer{}.start(cinema{})
	computer{}.start(micophone{})
}
