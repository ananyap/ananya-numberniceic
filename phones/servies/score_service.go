package servies

import "ananya/phones/utils"

type Score interface {
	PosSumScore() utils.PairPoint
	PosPairsAScore() []utils.PairPoint
	PosPairsBScore() []utils.PairPoint
	ConAScore() []utils.PairPoint
	ConBScore() []utils.PairPoint


}
