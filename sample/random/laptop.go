package random

import (
	"github.com/jasonrowsell/speccy/sample/constants"
	"github.com/jasonrowsell/speccy/utils"
)

func RandomLaptopBrand() string {
	return utils.RandomStringFromArray(constants.LaptopBrands)
}

func RandomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return utils.RandomStringFromArray(constants.AppleLaptopNames)
	case "Dell":
		return utils.RandomStringFromArray(constants.DellLaptopNames)
	case "HP":
		return utils.RandomStringFromArray(constants.HpLaptopNames)
	case "Lenovo":
		return utils.RandomStringFromArray(constants.LenovoLaptopNames)
	case "Microsoft":
		return utils.RandomStringFromArray(constants.MicrosoftLaptopNames)
	case "Samsung":
		return utils.RandomStringFromArray(constants.SamsungLaptopNames)
	default:
		return utils.RandomStringFromArray(constants.AppleLaptopNames)
	}
}
