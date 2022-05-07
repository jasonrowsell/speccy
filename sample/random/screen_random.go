package random

import (
	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample/constants"
)

func RandomScreenType() pb.Screen_ScreenType {
	return constants.ScreenTypes[RandomInt(0, len(constants.ScreenTypes)-1)]
}
