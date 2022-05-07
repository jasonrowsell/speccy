package constants

import "github.com/jasonrowsell/speccy/pb"

var ScreenTypes = []pb.Screen_ScreenType{
	pb.Screen_LCD,
	pb.Screen_CRT,
	pb.Screen_OLED,
	pb.Screen_IPS,
}
