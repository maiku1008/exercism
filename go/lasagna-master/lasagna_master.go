package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	var gramsOfPasta int = 0
	var litersOfSauce float64 = 0
	for _, layer := range layers {
		if layer == "noodles" {
			gramsOfPasta += 50
			continue
		}
		if layer == "sauce" {
			litersOfSauce += 0.2
			continue
		}
	}
	return gramsOfPasta, litersOfSauce
}

func AddSecretIngredient(friendList, myList []string) []string {
	secret := friendList[len(friendList)-1]
	myList[len(myList)-1] = secret
	return myList
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaled := make([]float64, len(amounts))
	factor := float64(portions) / 2.0
	for i, amount := range amounts {
		scaled[i] = amount * factor
	}
	return scaled
}
