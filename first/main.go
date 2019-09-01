package main

import (
	"encoding/json"
	"fmt"
	"github.com/jsxz/go-tutorial/first/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func indexView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello world`))
}
func listView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/list.html")
	w.Write(buf)
}
func addView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/add.html")
	w.Write(buf)
}
func editView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/edit.html")
	w.Write(buf)
}
func userAll(w http.ResponseWriter, r *http.Request) {
	mods, err := model.UserAll()
	if (err != nil) {
		log.Fatal(err)
	}
	buf, _ := json.Marshal(mods)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(buf)
}
func userAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//name:=r.FormValue("name")
	//get post 都是这样接收数据
	name:=r.Form.Get("name")
	num:=r.Form.Get("num")
	password:=r.Form.Get("password")
	logo:=r.Form.Get("logo")
	age:=r.Form.Get("name")
	a,_:=strconv.Atoi(age)
	model.UserAdd(name,num,password,logo,a)
}
func userUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id,_:=strconv.Atoi(r.Form.Get("id"))
	//get post 都是这样接收数据
	name:=r.Form.Get("name")
	num:=r.Form.Get("num")
	password:=r.Form.Get("password")
	logo:=r.Form.Get("logo")
	age:=r.Form.Get("name")
	a,_:=strconv.Atoi(age)
	model.UserUpdate(name,num,password,logo,a,id)
}
func userDel(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//name:=r.FormValue("name")
	//get post 都是这样接收数据

	id,_:=strconv.Atoi(r.Form.Get("id"))
	ok:=model.UserDel(id)
	if ok{
		w.Write([]byte("删除成功"))
	}else {
		w.Write([]byte("删除失败"))
	}
}
func userOne(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	fmt.Println(id)
	mod, _ := model.UserOne(id)
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(buf)
}
func oneView(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("views/one.html")
	w.Write(buf)
}
func main() {
	//http.Handle("/res/",http.FileServer(http.Dir("/res/")))
	http.Handle("/res/",http.StripPrefix("/res/",http.FileServer(http.Dir("res/"))))
	http.HandleFunc("/", indexView)
	http.HandleFunc("/list", listView)
	http.HandleFunc("/one", oneView)
	http.HandleFunc("/add", addView)
	http.HandleFunc("/edit", editView)
	http.HandleFunc("/api/user/all", userAll)
	http.HandleFunc("/api/user/one", userOne)
	http.HandleFunc("/api/user/add", userAdd)
	http.HandleFunc("/api/user/del", userDel)
	http.HandleFunc("/api/user/update", userUpdate)
	http.ListenAndServe(":80", nil)
	fmt.Println("run")
}
