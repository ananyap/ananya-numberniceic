package phones

import (
	"ananya/phones/servies"
	"ananya/phones/utils"
)

type PointEngine struct {
	pairs       servies.Pairs
	pairsNumber []utils.PairNumber
	pairsCont servies.PairContinuous

}

func NewPointEngine(pairs servies.Pairs, pairsNumber []utils.PairNumber, pairsCont servies.PairContinuous) *PointEngine {
	return &PointEngine{
		pairs,
		pairsNumber,
		pairsCont,
	}
}
func (pointEngine *PointEngine) ConAScore() []utils.PairPoint {
	var scorePairPoint []utils.PairPoint
	if pointEngine.pairsCont.PairsAContinue().Amount > 0 {
		for _ , pair := range pointEngine.pairsCont.PairsAContinue().AmType{
					switch pair.Type {
					case "D10":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point + 100,
						})
						break
					case "D8":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point + 80,
						})
						break
					case "D5":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point + 50,
						})
						break
					case "R10":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point - 100,
						})
						break
					case "R7":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point - 70,
						})
						break
					case "R5":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point - 50,
						})
						break
					}


				}
			}

	return scorePairPoint
}

func (pointEngine *PointEngine) ConBScore() []utils.PairPoint {
	var scorePairPoint []utils.PairPoint
	if pointEngine.pairsCont.PairsBContinue().Amount > 0 {
		for _, pair := range pointEngine.pairsCont.PairsBContinue().AmType{
					switch pair.Type {
					case "D10":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point + 100,
						})
					case "D8":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point + 80,
						})
					case "D5":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point + 50,
						})
					case "R10":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point - 100,
						})
					case "R7":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point - 70,
						})
					case "R5":
						scorePairPoint = append(scorePairPoint, utils.PairPoint{
							Number:    pair.Number,
							PointPair: pair.Point - 50,
						})
					}


				}
			}

	return scorePairPoint
}

func (pointEngine *PointEngine) PosPairsBScore() (scorePairsB []utils.PairPoint) {
	for position, pair := range pointEngine.pairs.PairsB() {
		for _, number := range pointEngine.pairsNumber {
			if pair == number.Number {
				switch position {
				case 0:
					scorePairsB = append(scorePairsB, getScorePairBPos0(number))
				case 1:
					scorePairsB = append(scorePairsB, getScorePairBPos1(number))
				case 2:
					scorePairsB = append(scorePairsB, getScorePairBPos2(number))
				case 3:
					scorePairsB = append(scorePairsB, getScorePairBPos3(number))
				case 4:
					scorePairsB = append(scorePairsB, getScorePairBPos4(number))
				case 5:
					scorePairsB = append(scorePairsB, getScorePairBPos5(number))
				}

				break
			}
		}
	}

	return scorePairsB
}
func getScorePairBPos0(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairBPos1(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairBPos2(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairBPos3(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairBPos4(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairBPos5(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}

func (pointEngine *PointEngine) PosPairsAScore() (scorePairsA []utils.PairPoint) {
	for position, pair := range pointEngine.pairs.PairsA() {
		for _, number := range pointEngine.pairsNumber {
			if pair == number.Number {
				switch position {
				case 0:
					scorePairsA = append(scorePairsA, getScorePairAPos0(number))
				case 1:
					scorePairsA = append(scorePairsA, getScorePairAPos1(number))
				case 2:
					scorePairsA = append(scorePairsA, getScorePairAPos2(number))
				case 3:
					scorePairsA = append(scorePairsA, getScorePairAPos3(number))
				case 4:
					scorePairsA = append(scorePairsA, getScorePairAPos4(number))
				}
				break
			}
		}
	}

	return scorePairsA
}
func getScorePairAPos4(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairAPos3(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairAPos2(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairAPos1(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}
func getScorePairAPos0(number utils.PairNumber) utils.PairPoint {
	score := 0
	switch number.Type {
	case "D10":
		score += number.Point + 100
	case "D8":
		score += number.Point + 80
	case "D5":
		score += number.Point + 50
	case "R10":
		score += number.Point - 100
	case "R7":
		score += number.Point - 70
	case "R5":
		score += number.Point - 50
	}

	return utils.PairPoint{
		Number:    number.Number,
		PointPair: score,
	}
}

func (pointEngine *PointEngine) PosSumScore() (pairPointSum utils.PairPoint) {
	pairSum := pointEngine.pairs.PairSum()
	for _, v := range pointEngine.pairsNumber {
		if pairSum == v.Number {
			return getScorePairSum(v)
		}
	}
	return pairPointSum
}

func getScorePairSum(v utils.PairNumber) utils.PairPoint {
	score := 0
	switch v.Type {
	case "D10":
		score += v.Point + 100
	case "D8":
		score += v.Point + 80
	case "D5":
		score += v.Point + 50
	case "R10":
		score += v.Point - 100
	case "R7":
		score += v.Point - 70
	case "R5":
		score += v.Point - 50
	}
	return utils.PairPoint{
		Number:    v.Number,
		PointPair: score,
	}
}
