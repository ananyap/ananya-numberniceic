package main

import (
	"ananya/DB"
	"ananya/phones"
	"ananya/phones/servies"
	"ananya/phones/utils"
	"database/sql"
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var pairs servies.Pairs
var pairsNumberMiracle []utils.PairNumber
var pairsColor servies.PairsColor
var score servies.Score
var pairContinue servies.PairContinuous
var reportScore phones.ReportScore

func main() {
	router := gin.Default()
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root: "views",
		Extension: ".html",
		Master: "layouts/master",
		DisableCache: true,
	})

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{
			"title": "Page title!",
		})
	})

	phoneHandler := []gin.HandlerFunc{func(c *gin.Context) {
		phoneNumber := c.Param("number")
		repo := DB.NewNumberRepo(db())
		pairs = phones.NewPhoneNumber(phoneNumber)

		pairsNumberMiracle = getPairsNumberMiracle(repo, pairs)
		pairContinue = phones.NewContinueEngine(pairs, pairsNumberMiracle)
		pairsColor = phones.NewColorEngine(pairsNumberMiracle, pairs)
		score = phones.NewPointEngine(pairs, pairsNumberMiracle, pairContinue)
		reportScore = phones.NewReportScore(score)




		c.JSON(200, gin.H{
			"Api": getApiPhone(phoneNumber, pairsColor, score, pairContinue, reportScore),
		})
	}}

	router.GET("/phone/:number", phoneHandler...)
	_ = router.Run(":8080")

}



func getApiPhone(number string, pairsColor servies.PairsColor, score servies.Score, continuous servies.PairContinuous, reportScore phones.ReportScore) utils.ApiPhone {
	reportScore.SetScoreByPosA()
	reportScore.SetScoreByPosB()
	reportScore.SetScoreConA()
	reportScore.SetScoreConB()
	scoreByPosA := reportScore.GetScoreByPosA()
	scoreByPosB := reportScore.GetScoreByPosB()
	scoreByConA := reportScore.GetScoreConA()
	scoreByConB := reportScore.GetScoreConB()


	return utils.ApiPhone{
		PhoneNumber:    number,
		PairsAColor:    pairsColor.PairsAColor(),
		PairsBColor:    pairsColor.PairsBColor(),
		PairSumColor:   pairsColor.PairSumColor(),
		PosSumPairSore: score.PosSumScore(),
		PosPairsASore:  score.PosPairsAScore(),
		SumScoreAPos: scoreByPosA,
		PosPairsBSore:  score.PosPairsBScore(),
		SumScoreBPos: scoreByPosB,
		ContinueA:      continuous.PairsAContinue(),
		ContinueB:      continuous.PairsBContinue(),
		ConsAScore: score.ConAScore(),
		ConsBScore: score.ConBScore(),
		SumScoreConA: scoreByConA,
		SumScoreConB: scoreByConB,
	}
}

func getPairsNumberMiracle(repo *DB.NumberRepo, pairs servies.Pairs) (pairsNumber []utils.PairNumber) {
	pairsUnique := pairs.PairsUnique(pairs.PairSum(), pairs.PairsA(), pairs.PairsB())
	return repo.Miracle(pairsUnique)
}

func db() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@IntelliP24@tcp(localhost:3306)/ananyadb?charset=utf8mb4,utf8")
	if err != nil {
		panic(err.Error())
	}
	return db
}
