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

func randomCPUBrand() string {
	if randomBoolean() {
		return "AMD"
	} else {
		return "Intel"
	}
}

func randomCPUName() string {
	brand := randomCPUBrand()

	if brand == "AMD" {
		return randomStringFromArray(amdCPUNames)
	} else {
		return randomStringFromArray(intelCPUNames)
	}
}

func randomSocket() string {
	switch randomInt(0, 2) {
	case 0:
		return "LGA1151"
	case 1:
		return "LGA1155"
	default:
		return "LGA1151"
	}
}
