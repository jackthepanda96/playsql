package activity

import (
	"database/sql"
	"errors"
	"log"
)

type ActivityN struct {
	ID         int
	Title      string
	CreateDate string
	Location   string
	Owner      int
}

type ActivtyMenuN struct {
	DB *sql.DB
}

func (am *ActivtyMenuN) Insert(newActivity ActivityN) (int, error) {
	insertActivityQry, err := am.DB.Prepare("INSERT INTO activities (title, location, due_date, owner) values (?,?,?,?)")
	if err != nil {
		log.Println("prepare insert activity ", err.Error())
		return 0, errors.New("prepare statement insert user error")
	}

	res, err := insertActivityQry.Exec(newActivity.Title, newActivity.Location, newActivity.CreateDate, newActivity.Owner)

	if err != nil {
		log.Println("insert activity ", err.Error())
		return 0, errors.New("insert activity error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert activity ", err.Error())
		return 0, errors.New("error setelah insert activity")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return 0, errors.New("no record")
	}

	id, _ := res.LastInsertId()

	return int(id), nil
}
