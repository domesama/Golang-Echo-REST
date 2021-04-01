package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ilyakaznacheev/cleanenv"
)

type MySqlConf struct {
	User string `env:"DB_USER" env-default:"OwO"`
	Pwd string `env:"DB_PWD" env-default:"OwO"`
	DriverName string `env:"DB_DRIVER_NAME" env-default:"mysql"`
	Port string `env:"DB_PORT" env-default:"3306"`
	Host string `env:"DB_HOST" env-default:"localhost"`
	Schema string `env:"DB_SCHEMA" env-default:"learning_go"`
}

func InitDB() (db *sql.DB, err error) {
	mySqlConf := &MySqlConf{}
	if err = cleanenv.ReadEnv(mySqlConf); err != nil{
		return nil,err
	}

	//"OwO:OwO@tcp(localhost)/learning_go"
	db, err = sql.Open( mySqlConf.DriverName,
		fmt.Sprintf("%s:%s@tcp(%s)/%s", mySqlConf.User,mySqlConf.Pwd,mySqlConf.Host,mySqlConf.Schema))
	if err != nil || db == nil{
		return nil,err
	}
	return db,nil
}

