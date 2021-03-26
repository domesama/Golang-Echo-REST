package services

import (
	"database/sql"
	"fmt"
	"github.com/domesama/Golang-Echo-REST/packages/structs"
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

func (w *AnimeWaifuService) SendWaifus(waifus structs.AnimeWaifu) error  {

	statement, err := w.db.Prepare(`INSERT INTO waifus (name,email,title,hair_color,age)VALUES (?,?,?,?,?) `)
	if err != nil{
		return err
	}
	res, err :=statement.Exec(waifus.Name, waifus.Email, waifus.Title, waifus.HairColor, waifus.Age)
	if err != nil{
		return err
	}
	fmt.Print("SQL results: ",fmt.Sprintf("%v"),res)
	return nil
}