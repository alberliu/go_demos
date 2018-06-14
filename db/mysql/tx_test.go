package mysql

import (
	"database/sql"
	"testing"
	"fmt"
	"runtime"
)

// Session 事务会话
type Session struct {
	DB *sql.DB
	Tx *sql.Tx
	Pc uintptr
}

func (s *Session) Begin() error {
	if s.Pc == 0 {
		var err error
		s.Tx, err = db.Begin()
		if err != nil {
			return err
		}

		pc, _, _, _ := runtime.Caller(1)
		s.Pc = pc
	}
	return nil
}

func (s *Session) Rollback() error {
	if s.Tx != nil {
		return s.Tx.Rollback()
	}
	return nil
}

func (s *Session) Commit() error {
	if s.Tx != nil {
		pc, _, _, _ := runtime.Caller(1)
		if s.Pc == pc{
			err:=s.Tx.Commit()
			if err!=nil{
				return err
			}
		}
	}
	return nil
}

func (s *Session) Exec(query string, args ...interface{}) (sql.Result, error) {
	if s.Tx != nil {
		return s.Tx.Exec(query, args...)
	}
	return s.DB.Exec(query, args...)
}

func (s *Session) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.Tx.QueryRow(query, args...)
}

// UserService 用户service
type UserService struct {
	Session
}

func (s *UserService) Insert() error {
	_, err := s.Exec("insert into user(number,name,ege,sex) values(?,?,?,?)", "1", "1", 1, "1")
	return err
}

func (s *UserService) Get() (*User, error) {
	row := db.QueryRow("SELECT number,name,ege,sex FROM user WHERE name = ?", "1")
	user := new(User)
	err := row.Scan(&user.number, &user.name, &user.ege, &user.sex)
	return user, err
}

func TestDo(t *testing.T) {
	s := UserService{}
	s.DB = db
	fmt.Println(s.Get())
	err := s.Insert()
	fmt.Println(err)
}

func TestDoTx(t *testing.T) {
	var err error
	s := UserService{}
	s.Tx, err = db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s.Get())
	err = s.Insert()
	if err != nil {
		fmt.Println(err)
	}
	s.Commit()
}

func TestTest(t *testing.T) {
	var a uintptr
	fmt.Println(a)
}
