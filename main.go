package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sjmudd/mysql_defaults_file"
	"log"
	"net/http"
)

var sqlStatus string

func main() {
	http.HandleFunc("/", serverstatus)
	http.ListenAndServe(":8080", nil)
}

func dbcheck() string {
	fmt.Println("MySQL health check invoked")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	//db, err := sql.Open("mysql", "root:testing@tcp(127.0.0.1:3306)/test")
	dbh, err := mysql_defaults_file.OpenUsingDefaultsFile("mysql", "", "performance_schema")

	// if there is an error opening the connection, handle it

	if err != nil {
		//panic(err.Error())
		log.Fatal(err)
	}

	// defer the close till after the main function has finished
	// executing

	defer dbh.Close()

	//defer db.Close()
	var version string

	err2 := dbh.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		//log.Fatal(err2)
		sqlStatus = "NOTOK"
	} else {
		sqlStatus = "OK"
	}

	return sqlStatus

	//fmt.Println(version)

}

func serverstatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	var status string
	status = dbcheck()
	if status == "OK" {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}
