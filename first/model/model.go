package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)
import _ "github.com/go-sql-driver/mysql"

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Open(`mysql`, `root:123456@tcp(127.0.0.1:3306)/go-tutorial?charset=utf8&parseTime=true`)
	if (err != nil) {
		panic("连接错误")
	}
	if err = DB.Ping(); err != nil {
		panic("运行错误")
	}
}

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name":"name" form:"name"`
	Logo     string `json:"logo":"logo" form:"logo"`
	Num      string `json:"num" form:"num"`
	Password string `json:"password" form:"password"`
	Age      int    `json:"age" form:"age"`
}

func UserAll() ([]User, error) {
	mods := make([]User, 0)
	err := DB.Select(&mods, "select * from `user`")
	return mods, err
}

func UserOne(id int)(User, error)  {
	log.Println("查询："+strconv.Itoa(id))
	//mod := User{}
	mods := make([]User, 0)
	err :=DB.Select(&mods,"select * from `user` where id=?",id)
	if(err!= nil) {
		log.Fatal(err)
	}
	return  mods[0] ,err
}
func UserAdd(name,num,password,logo string,age int) bool {
	_,err:=DB.Exec("insert into `user` (`name`,`num`,`password`,`logo`,`age`) values(?,?,?,?,?)",name,num,password,logo,age)
	if(err!= nil) {
		log.Fatal(err)
	}
	fmt.Println(name+"+"+password+"+"+logo+"+"+num)
	return true
}
func UserDel(id int)bool  {
	result,_:=DB.Exec("delete from `user` where id=?",id)
	affect,_:=result.RowsAffected()
	if affect==1{
		return true
	}
	return false
}
func UserUpdate(name,num,password,logo string,age,id int) bool  {
	result,_:=DB.Exec("update  `user` set `name`=?,`num`=?,`password`=?,`logo`=?,`age`=? where `id`=?",name,num,password,logo,age,id)
	affect,_:=result.RowsAffected()
	if affect==1{
		return true
	}
	return false
}