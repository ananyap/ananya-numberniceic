package DB

import (
	"ananya/phones/utils"
	"database/sql"
)

type NumberRepo struct {
	Db *sql.DB
}



func NewNumberRepo(db *sql.DB) *NumberRepo {
	return &NumberRepo{
		Db: db,
	}
}

func (numberRepo *NumberRepo) Miracle(uniqueNumber []string) (repoObj []utils.PairNumber) {


	var pairNumber utils.PairNumber
	var pairNumbers []utils.PairNumber

	rows, _ := numberRepo.Db.Query("SELECT pairnumber, pairtype, pairpoint FROM numbers")
	defer numberRepo.Db.Close()
	for rows.Next() {
		err := rows.Scan(&pairNumber.Number, &pairNumber.Type, &pairNumber.Point)
		if err != nil {
			panic(err)
		}
		if err = rows.Err(); err != nil {
			panic(err)
		}

		number := utils.PairNumber{Number: pairNumber.Number, Type: pairNumber.Type, Point: pairNumber.Point}
		pairNumbers = append(pairNumbers, number)

	}


	for _, unique := range uniqueNumber {

		for _, pair := range pairNumbers {

			if  unique == pair.Number {
				repoObj = append(repoObj, pair)
				break
			}
		}
	}


	return
}
