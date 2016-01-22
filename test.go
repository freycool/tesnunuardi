package main
import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	_"database/sql"
	"encoding/json"
)

	var (
		//DB *sqlx.DB
		//Err error
	)

	const (
		db_driver = "mysql"
		db_source = "pasar:pasar@tcp(192.168.0.104:3306)/pasar?charset=utf8&timeout=3s"
	)

	func CheckErr(Err error) {
	if Err != nil {
		panic(Err)
		}
	}

	func AddDB(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
		c.Next()
		}
	}

	type device struct  {
		Id int
		Device_name string
		Host string
		Port int
		Token string
		Description string
	}


	func NewDevice(db *sqlx.DB, name string, host string, port string, token string, description string){
		tx := db.MustBegin()

		tx.MustExec("INSERT INTO device_list (device_name, host, port, token, description) VALUES (?, ?, ?, ?, ?)", name, host, port, token, description)
		tx.Commit()
	}


	func regisDevice(c *gin.Context){

		db := c.MustGet("database").(*sqlx.DB)

		name := c.PostForm("name")
		host := c.PostForm("host")
		port := c.PostForm("port")
		token := c.PostForm("token")
		description := c.PostForm("description")

		fmt.Println("DB: ", db.Ping())
		fmt.Println("Name: ", name)
		fmt.Println("Host: ", host)
		fmt.Println("Port: ", port)
		fmt.Println("Token: ", token)
		fmt.Println("Description: ", description)
		//fmt.Print(name, host, port, token, description)
		//NewDevice(db, name, host, port, token, description)
		c.String(http.StatusOK, "Berhasil gan %s %s %s %s %s ", name, host, port, token, description)
		NewDevice(db, name, host, port, token, description)

		//c.JSON(http.StatusOK, NewDevice)


	}

	func outDevice(c *gin.Context) {
		fmt.Println("Try to print available device")

		ab := c.MustGet("database").(*sqlx.DB)

		dev := []device{}

		err := ab.Select(&dev, "SELECT id, device_name, host, port, token, description FROM device_list")

		if err != nil {
			fmt.Println("error: ", err)
		}

		//encode json
		jsn, err := json.Marshal(dev)


		fmt.Println("get data device: ", dev)
		fmt.Println("data to json :", string(jsn))

		//json to ajax
		c.JSON(http.StatusOK, dev)
		/*
		for i := range dev {
			fmt.Println("Device Name: ", dev[i].port)
		}*/
	}





func main() {
	DB, Err := sqlx.Open(db_driver, db_source)
	CheckErr(Err)
	if DB.Ping() != nil {
		fmt.Println("Cannot connect to Database")
	} else {
		fmt.Println("Connected to database")
	}

	//menyalakan web server
	r := gin.Default()


	r.Use(AddDB(DB))



	//load template

	r.LoadHTMLGlob("./template/*")
	r.GET("/testinput", func(c *gin.Context) {
		c.HTML(http.StatusOK, "testinput.tmpl", gin.H{

		})
	})

	r.POST("/coba", regisDevice)


	r.GET("/out", outDevice)




	//tx.MustExec("INSERT INTO device_list (device_name, host, port, token, description) VALUES ($1, $2, $3, $4, $5)", "arduino", "29", "33", "3", "tes masuk")



	r.Run(":8080") // listen and serve on 0.0.0.0:8080


}




