package appDB


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySQLDB struct   {
	url string
	db *sql.DB
}
//
var MySQL * MySQLDB

func Connect(url string){

	MySQL = new(MySQLDB)
	MySQL.url = url

	db, err := sql.Open("mysql", MySQL.url)
	if err != nil {
		log.Panicf("can not connect: %v", err)
	}
	MySQL.db = db
}

func (mysql *MySQLDB) Exec(query string) sql.Result{

	ret, err := mysql.db.Exec(query)
	if err != nil {
		log.Panicf("query failed: %v", err)
	}
	return ret
}

func (mysql *MySQLDB) Query(query string) *sql.Rows{

	rows, err := mysql.db.Query(query)
	if err != nil {
		log.Panicf("query failed: %v", err)
	}
	return rows
}

func (mysql *MySQLDB) Init(){

	_ = mysql.Exec("CREATE SCHEMA IF NOT EXISTS weather;")
	_ = mysql.Exec("USE weather")

	_ = mysql.Exec(`CREATE TABLE IF NOT EXISTS weather
(
city VARCHAR(24) NOT NULL,
weather VARCHAR(24) NOT NULL,
temperature FLOAT NOT NULL,
check_time TIMESTAMP DEFAULT NOW() null
);`)

}