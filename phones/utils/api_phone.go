package utils

type ApiPhone struct {
	PhoneNumber    string
	PosSumPairSore PairPoint
	PairSumColor   PairColor
	PosPairsASore  []PairPoint
	SumScoreAPos   ScoreDR
	PairsAColor    []PairColor
	ContinueA      ConType
	ConsAScore     []PairPoint
	SumScoreConA ScoreDR
	PosPairsBSore  []PairPoint
	PairsBColor    []PairColor
	SumScoreBPos   ScoreDR
	ContinueB      ConType
	ConsBScore     []PairPoint
	SumScoreConB ScoreDR
}
