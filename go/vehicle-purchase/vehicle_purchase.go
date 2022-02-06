package purchase

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	recommendation := ""
	if option1 < option2 {
		recommendation = option1
	}
	if option1 > option2 {
		recommendation = option2
	}
	if option1 == option2 {
		recommendation = "Both"
	}
	return fmt.Sprintf("%s is clearly the better choice.", recommendation)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var percentage float64 = 0
	switch {
	case age >= 3 && age < 10:
		percentage = 70
	case age >= 10:
		percentage = 50
	case age < 3:
		percentage = 80
	}
	return (originalPrice / 100) * percentage
}
