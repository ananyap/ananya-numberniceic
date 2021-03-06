package phones

import (
	"ananya/phones/servies"
	"ananya/phones/utils"
)

type ContinueEngine struct {
	pair        servies.Pairs
	pairsNumber []utils.PairNumber
}

func NewContinueEngine(pair servies.Pairs, number []utils.PairNumber) *ContinueEngine {
	return &ContinueEngine{
		pair:        pair,
		pairsNumber: number,
	}
}

func (contEngine *ContinueEngine) PairsAContinue() utils.ConType {
	amount := 0
	var typePairs []utils.PairNumber
	statusType := ""

	for i := 4; i >= 0; i-- {
		if i == 4 {
			count, t, amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsA()[i])
			amount += count
			typePairs = append(typePairs, utils.PairNumber{
				Number: amPair,
				Type:   t,
				Point:  point,
			})
			statusType = string([]rune(t)[0:1])

		}
		if i == 3 {
			count, t , amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsA()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})

			}else{
				break
			}
		}
		if i == 2 {
			count, t , amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsA()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})
			}else{
				break
			}
		}
		if i == 1 {
			count, t , amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsA()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})
			}else{
				break
			}
		}
		if i == 0 {
			count, t , amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsA()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})
			}else{
				break
			}
		}
	}

	return utils.ConType{
		Amount: amount,
		AmType: typePairs,
	}
}

func (contEngine *ContinueEngine) PairsBContinue() utils.ConType {
	amount := 0
	var typePairs []utils.PairNumber
	statusType := ""

	for i := 5; i >= 0; i-- {

		if i == 5 {
			count, t, amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsB()[i])
			amount += count
			statusType = string([]rune(t)[0:1])
			typePairs = append(typePairs, utils.PairNumber{
				Number: amPair,
				Type:   t,
				Point:  point,
			})

		}
		if i == 4 {
			count, t , amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsB()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})
			}else{
				break
			}


		}
		if i == 3 {
			count, t , amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsB()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})
			}else{
				break
			}
		}
		if i == 2 {
			count, t , amPair , point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsB()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})
			}else{
				break
			}
		}
		if i == 1 {
			count, t , amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsB()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})
			}else{
				break
			}
		}
		if i == 0 {
			count, t , amPair, point := getAmount(contEngine.pairsNumber, contEngine.pair.PairsB()[i])
			if statusType == string([]rune(t)[0:1]){
				amount += count
				typePairs = append(typePairs, utils.PairNumber{
					Number: amPair,
					Type:   t,
					Point:  point,
				})
			}else{
				break
			}
		}
	}

	return utils.ConType{
		Amount: amount,
		AmType: typePairs,
	}
}

func getAmount(pairsNumber []utils.PairNumber, pairs string) (amount int, amType string, amPair string, point int) {

	for _, numberObj := range pairsNumber {
		if pairs == numberObj.Number {

				if numberObj.Type == "D10" || numberObj.Type == "D8" || numberObj.Type == "D5" {
					amount += 1
					amType = numberObj.Type
					amPair = numberObj.Number
					point = numberObj.Point
					break
				}

				if numberObj.Type == "R10" || numberObj.Type == "R7" || numberObj.Type == "R5" {
					amount += 1
					amType = numberObj.Type
					amPair = numberObj.Number
					point = numberObj.Point
					break
				}
			}

		}


	return amount, amType, amPair, point
}
