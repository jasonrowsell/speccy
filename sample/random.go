package sample

import "github.com/jasonrowsell/speccy/pb"

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch randomInt(0, 3) {
	case 0:
		return pb.Keyboard_QWERTY
	case 1:
		return pb.Keyboard_QWERTZ
	case 2:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_QWERTY
	}
}

func randomBoolean() bool {
	return randomInt(0, 1) == 1
}
