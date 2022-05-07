package constants

import "github.com/jasonrowsell/speccy/pb"

var KeyboardLayouts = []pb.Keyboard_Layout{
	pb.Keyboard_QWERTY,
	pb.Keyboard_QWERTZ,
	pb.Keyboard_AZERTY,
}

var AmdCPUNames = []string{
	"AMD Athlon(tm) 64 X2 Dual Core Processor",
	"AMD Ryzen 5 3600",
	"AMD Ryzen 7 2700",
	"AMD Ryzen 7 2700X",
}

var IntelCPUNames = []string{
	"Intel Core i7-9700K",
	"Intel Core i7-9750H",
	"Intel Xeon E3-1280 V2",
	"Intel Xeon E3-1280 V3",
}

var SocketNames = []string{
	"LGA1151",
	"LGA1155",
	"LGA1156",
}
