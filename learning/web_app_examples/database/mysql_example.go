package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3309)/myapp?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT users SET _id=?,username=?,password=?,profession=?")
	checkErr(err)

	_id := time.Now().Unix()

	res, err := stmt.Exec(_id, "Saiham", "123", "Software Engineer")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// update
	stmt, err = db.Prepare("update users set username=? where _id=?")
	checkErr(err)

	res, err = stmt.Exec("Sayed", _id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	for rows.Next() {
		var id int
		var username string
		var password string
		var profession string

		err = rows.Scan(&id, &username, &password, &profession)
		checkErr(err)

		fmt.Println(id)
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(profession)
	}

	// delete
	stmt, err = db.Prepare("delete from users where _id=?")
	checkErr(err)

	res, err = stmt.Exec(_id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
