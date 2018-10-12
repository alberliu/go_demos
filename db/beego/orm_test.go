package beego

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id           int
	MobileNumber string
	Name         string
	Age          int
	Sex          int
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:Liu123456@tcp(localhost:3306)/test?charset=utf8", 30)
	// register model
	orm.RegisterModel(new(User))

	orm.Debug = true
}

func TestInsert(t *testing.T) {
	o := orm.NewOrm()
	user := User{
		Id:           103,
		MobileNumber: "",
		Name:         "1",
		Age:          1,
		Sex:          1,
	}

	id, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}

/*
func TestUpdate(t *testing.T) {
	o := orm.NewOrm()
	user := User{Id: 46, Number: "2", Name: "2", Ege: 2, Sex: 2, CreateTime: Time{}, UpdateTime: Time{}}
	num, err := o.Update(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}

func TestDelete(t *testing.T) {
	o := orm.NewOrm()
	user := User{Id: 45}
	num, err := o.Delete(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
*/
func TestGet(t *testing.T) {
	o := orm.NewOrm()
	user := User{Id: 1}
	err := o.Read(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", user)
	byt, _ := json.Marshal(user)
	fmt.Println(string(byt))
	var user1 User
	json.Unmarshal(byt, &user1)
	fmt.Printf("%+v", user1)

}

/*
func TestList(t *testing.T) {
	o := orm.NewOrm()
	qs := o.QueryTable(&User{})
	qs.Count()

	user := User{Id: 44}
	err := o.Read(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", user)
}

func TestSqlBuild(t *testing.T) {
	b, _ := orm.NewQueryBuilder("mysql")
	b.Select("*").
		From("user")
	o := orm.NewOrm()
	fmt.Println(b.String())
	var users []User
	_, err := o.Raw(b.String(), ).QueryRows(&users)
	if err != nil {
		fmt.Println(err)
	}
	printfln(users)
}

func TestSqlBuildCount(t *testing.T) {
	b, _ := orm.NewQueryBuilder("mysql")
	b.Select("count(*)").
		From("user")
	o := orm.NewOrm()
	var count int
	err := o.Raw(b.String()).QueryRow(&count)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
}

func printfln(users []User) {
	for _, user := range users {
		fmt.Printf("%+v", user)
	}
}

func TestTx(t *testing.T) {
	o := orm.NewOrm()
	o.Begin()
	defer o.Rollback()

	user1 := User{Id: 1, Number: "1", Name: "1", Ege: 1, Sex: 1, CreateTime: Time{}, UpdateTime: Time{}}
	_, err := o.Insert(&user1)
	if err != nil {
		fmt.Println(err)
	}

	user2 := User{Id: 2, Number: "2", Name: "2", Ege: 2, Sex: 2, CreateTime: Time{}, UpdateTime: Time{}}
	o.Insert(&user2)
	if err != nil {
		fmt.Println(err)
	}

	o.Commit()
}
*/
