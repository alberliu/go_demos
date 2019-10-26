package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:liu123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

type User struct {
	number string
	name   string
	ege    int
	sex    int
}

func TestExec(t *testing.T) {
	db.Exec("sql", 1, 2)

	stmt, err := db.Prepare(`insert into user(number,name,ege,sex) values(?,?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec("18829291353", "alber", 20, 1)
	checkErr(err)
	aff, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(aff)
}

func TestQuery(t *testing.T) {
	stmt, err := db.Prepare("select number,name,ege,sex from user")
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.number, &user.name, &user.ege, &user.sex)
		checkErr(err)
		fmt.Println(user)
		//rows.s
		col, err := rows.Columns()
		checkErr(err)
		fmt.Println(col)
	}
}

func TestTx(t *testing.T) {
	sql := `select
			number,
			name,
			ege,
			sex
		  from
		  	user`
	tx, err := db.Begin()
	stmt, err := tx.Prepare(sql)
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)
	users := make([]User, 0, 5)
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.number, &user.name, &user.ege, &user.sex)
		checkErr(err)
		fmt.Printf("%#v", user)
		users = append(users, user)
	}
	fmt.Println(users)
	tx.Commit()
	tx.Rollback()
}

func TestTxExec(t *testing.T) {
	tx, err := db.Begin()
	tx, err = db.Begin()
	stmt, err := tx.Prepare("insert into user(number,name,ege,sex) values(?,?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("18829291354", "alber", 20, 1)
	checkErr(err)
	aff, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(aff)
	tx.Commit()
}

func TestQueryRow(t *testing.T) {
	var age int64
	row := db.QueryRow("SELECT age FROM users WHERE name = ?", "alber")
	err := row.Scan(&age)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// 1866501296
func TestInsert1(t *testing.T) {
	defer duration(time.Now())
	for i := 0; i < 5000; i++ { //use b.N for looping
		_, err := db.Exec("insert into user1(name) values(?)", i)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// 1870954010
func TestInsert2(t *testing.T) {
	defer duration(time.Now())
	a := 0
	for i := 0; i < 5000; i++ { //use b.N for looping
		_, err := db.Exec("insert into user2(id,name) values(?,?)", a, a)
		if err != nil {
			fmt.Println(err)
		}
		a++
	}
}

// 3957588725
// 2892321216
func TestSelect(t *testing.T) {
	defer duration(time.Now())

	for i := 0; i < 50000; i++ { //use b.N for looping
		_ = db.QueryRow("select id,name from user1 where id = 1000")

	}
}

// 2892321216
func TestSelect2(t *testing.T) {
	defer duration(time.Now())

	for i := 0; i < 50000; i++ { //use b.N for looping
		_ = db.QueryRow("select id,name from user1 where name = 1000")

	}
}

func duration(t1 time.Time) {
	fmt.Println(time.Now().Sub(t1).Nanoseconds())
}
