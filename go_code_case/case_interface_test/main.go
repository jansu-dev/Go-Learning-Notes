package main

import "fmt"

type cinema struct{}

func (ca cinema) work() {
	fmt.Println("cinema is working!")
}

func (ca cinema) finish() {
	fmt.Println("cinema is finish!")
}

type micophone struct{}

func (com micophone) work() {
	fmt.Println("micophone is working!")
}

func (com micophone) finish() {
	fmt.Println("micophone is finish!")
}

type computer struct{}
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
