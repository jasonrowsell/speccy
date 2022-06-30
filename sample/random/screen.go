package random

import (
	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample/constants"
	"github.com/jasonrowsell/speccy/utils"
)

func RandomScreenType() pb.Screen_ScreenType {
	return constants.ScreenTypes[utils.RandomInt(0, len(constants.ScreenTypes)-1)]
}
