package sample

import (
	"github.com/jasonrowsell/speccy/pb"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:    randomKeyboardLayout(),
		Backlight: randomBoolean(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {
	name := randomCPUName()
	minGhz := randomFloat64(1, 3)
	maxGhz := randomFloat64(minGhz, 5)

	cpu := &pb.CPU{
		Name:       name,
		Cores:      uint32(randomInt(1, 4)),
		ClockSpeed: uint32(randomInt(1, 4) * 1000),
		Threads:    uint32(randomInt(1, 4)),
		Socket:     randomSocket(),
		CacheSize:  uint32(randomInt(1, 4) * 1024),
		MinGHz:     minGhz,
		MaxGHz:     maxGhz,
	}
	return cpu
}
