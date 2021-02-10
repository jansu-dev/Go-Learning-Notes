package main

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
	computer.start(cinema)
}
