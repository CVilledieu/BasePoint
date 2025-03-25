package main

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func main() {
	title := "Hello world"
	message := "I am a message"
	t, _ := syscall.UTF16PtrFromString(title)
	txt, _ := syscall.UTF16PtrFromString(message)

	windows.MessageBox(0, txt, t, 0)
}
