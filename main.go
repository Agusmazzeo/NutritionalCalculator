package main

import (
	"fmt"
)

func main() {

	ns := GetNutritionalScore(
		NutritionalData{
			Energy:              EnergyFromKcal(500),
			Sugars:              SugarGram(20),
			SaturatedFattyAcids: SaturatedFattyAcidsGram(2),
			Sodium:              SodiumMilligram(500),
			Fruits:              FruitsPercent(60),
			Fibre:               FibreGram(4),
			Protein:             ProteinGram(2),
		}, Food)

	fmt.Printf("Nutritional Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s \n", ns.GetNutriScore())
}
