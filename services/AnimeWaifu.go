package services

import (
	"database/sql"
	"errors"
	"github.com/domesama/Golang-Echo-REST/packages/structs"
	_ "github.com/go-sql-driver/mysql"
)

type AnimeWaifuService struct {
	db *sql.DB
}

func NewAnimeWaifuService(db *sql.DB) *AnimeWaifuService{
	statement, _ := db.Prepare(`
		create table if not exists waifus
		(
			id int auto_increment,
			name varchar(500) null,
		    email varchar(500) null,
		    title varchar(500) null,
		    hair_color varchar(500) null,
		    age int null,
			constraint waifus_pk
			primary key (id)
		);
	`)
	statement.Exec()
	return &AnimeWaifuService{
		db: db,
	}
}

func (w *AnimeWaifuService) GetWaifus() (waifus []structs.AnimeWaifu,err error){

	var id int
	var name string
	var email string
	var title string
	var hairColor string
	var age int

	rows, err := w.db.Query(`SElECT * FROM waifus`)
	if err != nil{
		return waifus,err
	}
	defer rows.Close()

	for rows.Next(){
		if err = rows.Scan(&id, &name, &email, &title, &hairColor, &age); err != nil{
			return waifus,err
		}
		waifus = append(waifus,
			structs.AnimeWaifu{
				Id: id, Name: name, Email: email, Title: title, HairColor: hairColor, Age: age})
	}

	if err = rows.Err(); err != nil {
		return waifus,err
	}

	return waifus,nil
}

func (w *AnimeWaifuService) SendWaifus(waifus structs.AnimeWaifu) error  {
	statement, err := w.db.Prepare(`INSERT INTO waifus (name,email,title,hair_color,age)VALUES (?,?,?,?,?) `)
	if err != nil{
		return err
	}
	_, err =statement.Exec(waifus.Name, waifus.Email, waifus.Title, waifus.HairColor, waifus.Age)
	if err != nil{
		return err
	}
	//fmt.Print("SQL results: ",fmt.Sprintf("%v"),res)
	return err
}

func (w *AnimeWaifuService) DeleteWaifu(id int) (err error) {
	statement, err := w.db.Prepare(`DELETE FROM waifus WHERE id=?`)

	row, err := statement.Exec(id)

	if err != nil{
		return err
	}
	affected,_ := row.RowsAffected()
	if affected <= int64(0) {
		return errors.New("There was no new rows affected")
	}
	return nil
}

