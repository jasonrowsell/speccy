package sample

import (
	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample/random"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:    random.RandomKeyboardLayout(),
		Backlight: random.RandomBoolean(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {
	name := random.RandomCPUName()
	minGhz := random.RandomFloat64(1, 3)
	maxGhz := random.RandomFloat64(minGhz, 5)

	cpu := &pb.CPU{
		Name:       name,
		Cores:      uint32(random.RandomInt(1, 4)),
		ClockSpeed: uint32(random.RandomInt(1, 4) * 1000),
		Threads:    uint32(random.RandomInt(1, 4)),
		Socket:     random.RandomSocket(),
		CacheSize:  uint32(random.RandomInt(1, 4) * 1024),
		MinGHz:     minGhz,
		MaxGHz:     maxGhz,
	}
	return cpu
}

func NewGPU() *pb.GPU {
	name := random.RandomGPUName()
	busId := random.RandomBusID()
	deviceId := random.RandomDeviceID()
	memory := &pb.Memory{
		Size: uint64(random.RandomInt(1, 4) * 1024),
		Unit: random.RandomMemoryUnit(),
	}

	gpu := &pb.GPU{
		Name:             name,
		BusId:            busId,
		DeviceId:         deviceId,
		Memory:           memory,
		ClockSpeed:       uint32(random.RandomInt(1, 4) * 1000),
		MemoryClockSpeed: uint32(random.RandomInt(1, 4) * 1000),
	}
	return gpu
}
