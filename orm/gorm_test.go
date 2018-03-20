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
	db, err := gorm.Open("mysql", "smartwatch:smartwatch_secret@tcp(106.75.177.192:13331)/test?charset=utf8")
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
