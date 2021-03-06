package servies

import "ananya/phones/utils"

type Pairs interface {
	PairsA() []string
	PairsB() []string
	PairSum() string
	PairsUnique(pairSum string, pairsA []string, pairsB []string) []string
}

type PairsColor interface {
	PairsAColor() []utils.PairColor
	PairsBColor() []utils.PairColor
	PairSumColor() utils.PairColor
}

type PairContinuous interface {
	PairsAContinue() utils.ConType
	PairsBContinue() utils.ConType
	//PairsRAContinue() int
	//PairsRBContinue() int
}