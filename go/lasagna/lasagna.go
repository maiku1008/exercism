package lasagna

// OvenTime returns how many minutes the lasagna should be in the oven
func OvenTime() int {
	return 40
}

// RemainingOvenTime returns how many minutes the lasagna still has to remain in the oven
func RemainingOvenTime(remaining int) int {
	return OvenTime() - remaining
}

// PreparationTime() returns how many minutes you spent preparing the lasagna, assuming each layer takes you 2 minutes to prepare
func PreparationTime(layers int) int {
	return 2 * layers
}

// ElapsedTime returns the total elapsed working time
func ElapsedTime(layers, minutes int) int {
	return PreparationTime(layers) + OvenTime() - RemainingOvenTime(minutes)
}
