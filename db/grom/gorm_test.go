package grom

import (
	"testing"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type User struct {
	Id     int64  `gorm:"primary_key"`
	Number string `gorm:"default":'alber'`
	Name   string
	Ege    int
	Sex    int
	Books  []Book `gorm:"ForeignKey:user_id;AssociationForeignKey:Refer"`
}

type Book struct {
	Id   int64 `gorm:"primary_key"`
	Name string
}

func initGORMDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Liu123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)
	db.SingularTable(true)
	return db;
}

func TestInsert(t *testing.T) {
	db := initGORMDB()
	user1 := User{Id: 108, Number: "", Name: "1", Ege: 1, Sex: 1}
	db = db.Create(&user1)

	user2 := User{Id: 109, Number: "", Name: "1", Ege: 1, Sex: 1}
	db = db.Create(&user2)
	fmt.Println(db.Error)
	fmt.Println(db.RowsAffected)
}

func TestInsertOneToMany(t *testing.T) {
	db := initGORMDB()
	user := User{Id: 0, Number: "1", Name: "1", Ege: 0, Sex: 1}

	books := make([]Book, 3)
	books[0] = Book{Name: "1"}
	books[1] = Book{Name: "2"}
	books[2] = Book{Name: "3"}

	user.Books = books
	db = db.Set("gorm:save_associations", true).Create(&user)
	//db = db.Create(&user)
	fmt.Println(db.Error)
	fmt.Println("hello world")
	fmt.Println(db.RowsAffected)

}

func TestGet(t *testing.T) {
	db := initGORMDB()
	user := User{}
	err:=db.First(&user, 10).Error
	if err!=nil{
		fmt.Println("err")
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user)
}

func TestRows(t *testing.T) {
	db := initGORMDB()
	rows, err := db.Raw("SELECT id FROM pharmacist_goods WHERE id = ?", 10000).Rows()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows.Columns())

	for rows.Next() {
		var id int
		rows.Scan(&id)
		fmt.Println("id:", id)
	}
}

func TestRow(t *testing.T) {
	db := initGORMDB()
	var id int
	err := db.Raw("SELECT id FROM pharmacist_goods WHERE id = ?", 1).Row().Scan(&id)
	if err != nil {
		fmt.Println(err)
	}

}
