package constants

import "github.com/jasonrowsell/speccy/pb"

var GPUBrands = []string{
	"Nvidia",
	"AMD",
	"Intel",
}

var NvidiaGPUNames = []string{
	"GeForce GTX 1080 Ti",
	"GeForce GTX 1080",
	"GeForce RTX 2080 Ti",
	"GeForce RTX 2080",
}

var AmdGPUNames = []string{
	"RX Vega 56",
	"RX Vega 64",
	"RX 580",
	"IGP-10",
}

var IntelGPUNames = []string{
	"Intel UHD Graphics 630",
	"Intel Iris Plus Graphics 640",
	"Intel HD Graphics 630",
	"Intel Iris Plus Graphics 650",
}

var MemoryUnits = []pb.Memory_Unit{
	pb.Memory_BYTES,
	pb.Memory_KILOBYTES,
	pb.Memory_MEGABYTES,
	pb.Memory_GIGABYTES,
	pb.Memory_TERABYTES,
}
