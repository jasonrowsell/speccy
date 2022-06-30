package random

import (
	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample/constants"
	"github.com/jasonrowsell/speccy/utils"
)

func randomGPUBrand() string {
	return utils.RandomStringFromArray(constants.GPUBrands)
}

func RandomGPUName() string {
	brand := randomGPUBrand()

	switch brand {
	case "Nvidia":
		return utils.RandomStringFromArray(constants.NvidiaGPUNames)
	case "AMD":
		return utils.RandomStringFromArray(constants.AmdGPUNames)
	case "Intel":
		return utils.RandomStringFromArray(constants.IntelGPUNames)
	default:
		return utils.RandomStringFromArray(constants.NvidiaGPUNames)
	}
}

func RandomBusID() string {
	id := utils.GenerateUUID()
	return id[:8]
}

func RandomDeviceID() string {
	id := utils.GenerateUUID()
	return id[:12]
}

func RandomMemoryUnit() pb.Memory_Unit {
	return constants.MemoryUnits[utils.RandomInt(0, len(constants.MemoryUnits)-1)]
}
