package phones

import (
	"ananya/phones/servies"
	"ananya/phones/utils"
)

type ReportScore struct {
	score servies.Score
	scoreByPosA utils.ScoreDR
	scoreByPosB utils.ScoreDR
	scoreByConA utils.ScoreDR
	scoreByConB utils.ScoreDR
}

func NewReportScore(score servies.Score) ReportScore {
	return ReportScore{
		score: score,
	}
}

func (r *ReportScore) SetScoreConB() {
	var pointD = 0
	var pointR = 0
	for _, s := range r.score.ConBScore(){
		if s.PointPair >= 0 {
			pointD += s.PointPair
		}else{
			pointR += s.PointPair
		}
	}
	r.scoreByConB = utils.ScoreDR{
		ScoreD: pointD,
		ScoreR: pointR,
	}
}

func (r *ReportScore) GetScoreConB() utils.ScoreDR {
	return r.scoreByConB
}

func (r *ReportScore) SetScoreConA() {
	var pointD = 0
	var pointR = 0
	for _, s := range r.score.ConAScore(){
		if s.PointPair >= 0 {
			pointD += s.PointPair
		}else{
			pointR += s.PointPair
		}
	}
	r.scoreByConA = utils.ScoreDR{
		ScoreD: pointD,
		ScoreR: pointR,
	}
}

func (r *ReportScore) GetScoreConA() utils.ScoreDR {
	return r.scoreByConA
}

func (r *ReportScore) SetScoreByPosA() {
	var pointD = 0
	var pointR = 0
	for _, s := range r.score.PosPairsAScore() {
		if s.PointPair >= 0 {
			pointD += s.PointPair
		}else{
			pointR += s.PointPair
		}
	}

	r.scoreByPosA = utils.ScoreDR{
		ScoreD: pointD,
		ScoreR: pointR,
	}
}

func (r *ReportScore) GetScoreByPosA() utils.ScoreDR {
	return r.scoreByPosA
}

func (r *ReportScore) SetScoreByPosB() {
	var pointD = 0
	var pointR = 0
	for _, s := range r.score.PosPairsBScore() {
		if s.PointPair >= 0 {
			pointD += s.PointPair
		}else{
			pointR += s.PointPair
		}
	}

	r.scoreByPosB = utils.ScoreDR{
		ScoreD: pointD,
		ScoreR: pointR,
	}
}

func (r *ReportScore) GetScoreByPosB() utils.ScoreDR {
	return r.scoreByPosB
}