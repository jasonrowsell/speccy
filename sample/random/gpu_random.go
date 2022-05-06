package random

import (
	"github.com/google/uuid"
	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample/constants"
)

func randomGPUBrand() string {
	return RandomStringFromArray(constants.GPUBrands)
}

func RandomGPUName() string {
	brand := randomGPUBrand()

	switch brand {
	case "Nvidia":
		return RandomStringFromArray(constants.NvidiaGPUNames)
	case "AMD":
		return RandomStringFromArray(constants.AmdGPUNames)
	case "Intel":
		return RandomStringFromArray(constants.IntelGPUNames)
	default:
		return RandomStringFromArray(constants.NvidiaGPUNames)
	}
}

func RandomBusID() string {
	id := uuid.New().String()
	return id[:8]
}

func RandomDeviceID() string {
	id := uuid.New().String()
	return id[:12]
}

func RandomMemoryUnit() pb.Memory_Unit {
	return constants.MemoryUnits[RandomInt(0, len(constants.MemoryUnits)-1)]
}
