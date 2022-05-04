package sample

import "github.com/jasonrowsell/speccy/pb"

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:    randomKeyboardLayout(),
		Backlight: randomBoolean(),
	}
	return keyboard
}
