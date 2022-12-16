package activity

import (
	"database/sql"
	"errors"
	"log"
)

type Activity struct {
	id          int
	title       string
	create_date string
	location    string
	owner       int
}

// SET digunakan untuk mengisikan value kedalam properti
func (a *Activity) SetID(newID int) {
	a.id = newID
}
func (a *Activity) SetTitle(newTitle string) {
	a.title = newTitle
}
func (a *Activity) SetCreateDate(newDate string) {
	a.create_date = newDate
}
func (a *Activity) SetLocation(newLocation string) {
	a.location = newLocation
}
func (a *Activity) SetOwner(owner int) {
	a.owner = owner
}

// GET Membaca value properti
func (a *Activity) GetID() int {
	return a.id
}
func (a *Activity) GetTitle() string {
	return a.title
}
func (a *Activity) GetCreateDate() string {
	return a.create_date
}
func (a *Activity) GetLocation() string {
	return a.location
}
func (a *Activity) GetOwner() int {
	return a.owner
}

type ActivtyMenu struct {
	db *sql.DB
}

type ActivityInterface interface {
	Insert(newActivity Activity) (int, error)
}

// Dependency Injection
func NewActivityMenu(conn *sql.DB) ActivityInterface {
	return &ActivtyMenu{
		db: conn,
	}
}

func (am *ActivtyMenu) Insert(newActivity Activity) (int, error) {
	insertActivityQry, err := am.db.Prepare("INSERT INTO activities (title, location, due_date, owner) values (?,?,?,?)")
	if err != nil {
		log.Println("prepare insert activity ", err.Error())
		return 0, errors.New("prepare statement insert user error")
	}

	res, err := insertActivityQry.Exec(newActivity.GetTitle(), newActivity.GetLocation(), newActivity.GetCreateDate(), newActivity.GetOwner())

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
