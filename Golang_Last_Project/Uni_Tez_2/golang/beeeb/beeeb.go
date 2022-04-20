package beeeb

import (
	"github.com/gen2brain/beeep"
	"main.go/checkError"
)

func BebMessage(Title string, Message string) {

	err := beeep.Alert(Title, Message, "pn/ggo.icon")
	checkError.ErrorContr(err, "9")
}
