package random

import (
	"github.com/jasonrowsell/speccy/sample/constants"
)

func RandomLaptopBrand() string {
	return RandomStringFromArray(constants.LaptopBrands)
}

func RandomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return RandomStringFromArray(constants.AppleLaptopNames)
	case "Dell":
		return RandomStringFromArray(constants.DellLaptopNames)
	case "HP":
		return RandomStringFromArray(constants.HpLaptopNames)
	case "Lenovo":
		return RandomStringFromArray(constants.LenovoLaptopNames)
	case "Microsoft":
		return RandomStringFromArray(constants.MicrosoftLaptopNames)
	case "Samsung":
		return RandomStringFromArray(constants.SamsungLaptopNames)
	default:
		return RandomStringFromArray(constants.AppleLaptopNames)
	}
}
