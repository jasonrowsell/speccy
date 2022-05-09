package sample

import (
	"math"

	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample/random"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	minGhz := math.Floor(random.RandomFloat64(1, 3)*10) / 10
	maxGhz := math.Floor(random.RandomFloat64(minGhz, 5)*10) / 10

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
		Size: uint64(random.RandomInt(2, 6)),
		Unit: pb.Memory_GIGABYTES,
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

func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Size: uint64(random.RandomInt(4, 64)),
		Unit: pb.Memory_GIGABYTES,
	}

	return ram
}

func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Size: uint64(random.RandomInt(128, 1024)),
			Unit: pb.Memory_GIGABYTES,
		},
	}

	return ssd
}

func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Size: uint64(random.RandomInt(1, 6)),
			Unit: pb.Memory_TERABYTES,
		},
	}

	return hdd
}

func NewScreen() *pb.Screen {
	height := uint32(random.RandomInt(1, 4))
	width := height * 16 / 9

	screen := &pb.Screen{
		Resolution: &pb.Screen_Resolution{
			Height: height,
			Width:  width,
		},
		Size:           math.Floor(random.RandomFloat64(13, 17)*100) / 100,
		Type:           random.RandomScreenType(),
		HasTouchscreen: random.RandomBoolean(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {
	id := random.GenerateUUID()
	brand := random.RandomLaptopBrand()
	name := random.RandomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       id,
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRam(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storage:  []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightGrams{
			WeightGrams: uint32(random.RandomInt(1, 4) * 1000),
		},
		Price:       math.Floor(random.RandomFloat64(100, 1000)*100) / 100,
		ReleaseDate: timestamppb.Now(),
	}

	return laptop
}
