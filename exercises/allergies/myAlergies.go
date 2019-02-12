package allergies

import (
	"fmt"
	"math"
)

var allAlergens = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies2(alergens uint) (results []string) {
	fmt.Printf("Alergens: %v\n", alergens)
	if alergens == 0 {
		return results
	}
	for y := 10; y >= 0; y-- {
		num := uint(math.Pow(2, float64(y)))
		if alergens >= num {
			fmt.Printf("Larges match of: %v in %v\n", num, alergens)
			results = append(results, allAlergens[y])
			alergens -= uint(num)
		}
	}
	return results
}
