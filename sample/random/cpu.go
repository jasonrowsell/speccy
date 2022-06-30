package random

import (
	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample/constants"
	"github.com/jasonrowsell/speccy/utils"
)

func RandomKeyboardLayout() pb.Keyboard_Layout {
	return constants.KeyboardLayouts[utils.RandomInt(0, len(constants.KeyboardLayouts)-1)]
}

func randomCPUBrand() string {
	if utils.RandomBoolean() {
		return "AMD"
	} else {
		return "Intel"
	}
}

func RandomCPUName() string {
	brand := randomCPUBrand()

	if brand == "AMD" {
		return utils.RandomStringFromArray(constants.AmdCPUNames)
	} else {
		return utils.RandomStringFromArray(constants.IntelCPUNames)
	}
}

func RandomSocket() string {
	return utils.RandomStringFromArray(constants.SocketNames)
}
