package beego

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"testing"
	"fmt"
	"time"
)

type User struct {
	Id         int
	Number     string
	Name       string
	Ege        int
	Sex        int
	CreateTime time.Time
	UpdateTime time.Time
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
	user := User{Id: 0, Number: "", Name: "1", Ege: 1, Sex: 1, CreateTime: time.Now(), UpdateTime: time.Now()}
	id, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}

func TestUpdate(t *testing.T) {
	o := orm.NewOrm()
	user := User{Id: 46, Number: "2", Name: "2", Ege: 2, Sex: 2,  UpdateTime: time.Now()}
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

func TestGet(t *testing.T) {
	o := orm.NewOrm()
	user := User{Id: 46}
	err := o.Read(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user)
}

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
	fmt.Printf("%+v\n", user)
}

func TestSqlBuild(t *testing.T){
	b, _ := orm.NewQueryBuilder("mysql")
	b.Select("*").
		From("user")
	o := orm.NewOrm()
	fmt.Println(b.String())
	var users []User
	_,err:=o.Raw(b.String(), ).QueryRows(&users)
	if err!=nil{
		fmt.Println(err)
	}
	printfln(users)
}

func TestSqlBuildCount(t *testing.T){
	b, _ := orm.NewQueryBuilder("mysql")
	b.Select("count(*)").
		From("user")
	o := orm.NewOrm()
	var count int
	err:=o.Raw(b.String() ).QueryRow(&count)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(count)
}

func printfln(users []User){
	for _,user:=range users{
		fmt.Printf("%+v\n",user)
	}
}