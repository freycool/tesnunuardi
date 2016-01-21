package main
import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

	var (
		DB *sql.DB
		Err error
	)

	const (
		db_driver = "mysql"
		db_source = "root:root(192.168.0.105:3306)/pasar?charset=utf8&timeout=3s"
	)

	func initDb() {
		DB, Err = sql.Open(db_driver, db_source)
}

func main() {
}
