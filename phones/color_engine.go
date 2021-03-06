package phones

import (
	"ananya/phones/servies"
	"ananya/phones/utils"
)



type ColorEngine struct {
	pairs             servies.Pairs
	uniquePairsNumber []utils.PairNumber
}

func NewColorEngine(uniquePairsNumber []utils.PairNumber, pairs servies.Pairs) *ColorEngine {
	return &ColorEngine{
		pairs:             pairs,
		uniquePairsNumber: uniquePairsNumber,

	}
}


//implements service
func (manager *ColorEngine) PairsAColor() []utils.PairColor {

	getColorPair := func(pairType string) string {
		typeRD := []rune(pairType)
		t := string(typeRD[0:1])
		if t == "D" {
			return "GREEN"
		} else {
			return "RED"
		}
	}

	var pairsColor []utils.PairColor
	for _, pairA := range manager.pairs.PairsA() {
		for _, pairUniqueObj := range manager.uniquePairsNumber {
			if pairUniqueObj.Number == pairA {
				pairsColor = append(pairsColor, utils.PairColor{
					Pair:  pairA,
					Color: getColorPair(pairUniqueObj.Type),
				})
			}

		}

	}

	return pairsColor

}

func (manager *ColorEngine) PairsBColor() []utils.PairColor {

	getColorPair := func(pairType string) string {
		typeRD := []rune(pairType)
		t := string(typeRD[0:1])
		if t == "D" {
			return "GREEN"
		} else {
			return "RED"
		}
	}

	var pairsColor []utils.PairColor
	for _, pairB := range manager.pairs.PairsB() {
		for _, pairUniqueObj := range manager.uniquePairsNumber {
			if pairUniqueObj.Number == pairB {
				pairsColor = append(pairsColor, utils.PairColor{
					Pair:  pairB,
					Color: getColorPair(pairUniqueObj.Type),
				})
			}

		}

	}

	return pairsColor

}

func (manager *ColorEngine) PairSumColor() utils.PairColor {
	getColorPair := func(pairType string) string {
		typeRD := []rune(pairType)
		t := string(typeRD[0:1])
		if t == "D" {
			return "GREEN"
		} else {
			return "RED"
		}
	}

	var pairColor utils.PairColor

		for _, pairUniqueObj := range manager.uniquePairsNumber {
			if pairUniqueObj.Number == manager.pairs.PairSum() {
				pairColor = utils.PairColor{
					Pair:  pairUniqueObj.Number,
					Color: getColorPair(pairUniqueObj.Type),
				}
			}

		}



	return pairColor

}

