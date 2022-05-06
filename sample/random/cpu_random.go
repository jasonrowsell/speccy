package random

import (
	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample/constants"
)

func RandomKeyboardLayout() pb.Keyboard_Layout {
	return constants.KeyboardLayouts[RandomInt(0, len(constants.KeyboardLayouts)-1)]
}

func randomCPUBrand() string {
	if RandomBoolean() {
		return "AMD"
	} else {
		return "Intel"
	}
}

func RandomCPUName() string {
	brand := randomCPUBrand()

	if brand == "AMD" {
		return RandomStringFromArray(constants.AmdCPUNames)
	} else {
		return RandomStringFromArray(constants.IntelCPUNames)
	}
}

func RandomSocket() string {
	return RandomStringFromArray(constants.SocketNames)
}
