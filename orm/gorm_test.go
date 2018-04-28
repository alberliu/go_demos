package orm

import (
	"testing"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type User struct {
	Id     int64 `gorm:"primary_key"`
	Number string
	Name   string
	Ege    int
	Sex    int
}

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "yaoshitong:12345678@tcp(192.168.40.12:3306)/yaoshitong?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	db.SingularTable(true)
	return db;
}

func TestInsert(t *testing.T) {
	db := initDB()
	user := User{Id: 4, Number: "1", Name: "1", Ege: 1, Sex: 1}
	db = db.Create(&user)
	fmt.Println(db.Error)
	fmt.Println("hello world")
	fmt.Println(db.RowsAffected)

}

func TestGet(t *testing.T) {
	db := initDB()
	user := User{}
	db.First(&user, 31)
	fmt.Printf("%+v", user)
}

func TestRows(t *testing.T){
	db := initDB()
	rows,err:=db.Raw("SELECT id FROM pharmacist_goods WHERE id = ?",10000).Rows()
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(rows.Columns())

	for rows.Next(){
		var id int
		rows.Scan(&id)
		fmt.Println("id:",id)
	}
}

func TestRow(t *testing.T){
	db := initDB()
	var id int
	err:=db.Raw("SELECT id FROM pharmacist_goods WHERE id = ?",1).Row().Scan(&id)
	if err!=nil{
		fmt.Println(err)
	}



}