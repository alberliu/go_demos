package xorm

import (
	"testing"

	"fmt"
	"github.com/xormplus/xorm"
)

func engine()*xorm.Engine{
	engine, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err!=nil{
		fmt.Println(err)
	}
	return engine
}

func TestTransaction(t *testing.T){
	session := engine().NewSession()
	defer session.Close()
	// add BeginTrans() before any action
	tx, err := session.BeginTrans()
	if err != nil {
		return
	}

	_, err = tx.Session().Insert(&user1)
	if err != nil {
		tx.Rollback()
		return
	}


		tx.RollbackTrans()


	// add CommitTrans() after all actions
	err = tx.CommitTrans()

}