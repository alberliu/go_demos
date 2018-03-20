package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"testing"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "smartwatch:smartwatch_secret@tcp(106.75.177.192:13331)/test?charset=utf8")
}

type User struct {
	number string
	name   string
	ege    int
	sex    int
}

func TestExec(t *testing.T) {
	stmt, err := db.Prepare("insert into user(number,name,ege,sex) values(?,?,?,?)")
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
	rows,err:=stmt.Query()
	checkErr(err)
	for rows.Next() {
		user:=User{}
		err = rows.Scan(&user.number, &user.name, &user.ege,&user.sex)
		checkErr(err)
		fmt.Println(user)
		//rows.s
		col,err:=rows.Columns()
		checkErr(err)
		fmt.Println(col)

	}


}

func TestTx(t *testing.T) {
	sql:=`select
			number,
			name,
			ege,
			sex
		  from
		  	user`
	tx,err:=db.Begin()
	stmt, err := tx.Prepare(sql)
	checkErr(err)
	rows,err:=stmt.Query()
	checkErr(err)
	users:=make([]User,0,5)
	for rows.Next() {
		user:=User{}
		err = rows.Scan(&user.number, &user.name, &user.ege,&user.sex)
		checkErr(err)
		fmt.Printf("%#v\n",user)
		users=append(users,user)
	}
	fmt.Println(users)
	tx.Commit()
}

func TestTxExec(t *testing.T) {
	tx,err:=db.Begin()
	stmt, err := tx.Prepare("insert into user(number,name,ege,sex) values(?,?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("18829291354", "alber", 20, 1)
	checkErr(err)
	aff, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(aff)
	tx.Commit()
}


func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

