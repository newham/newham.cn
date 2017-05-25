package controllers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	name     string
	password string
}
type Json struct {
	Title   string
	Content string
	User    User
}

/*
*return a json file
 */
func JsonController(w http.ResponseWriter, r *http.Request) {
	//有两种方式构造json数据，一种是struct，一种是map
	//两种方式的区别在于，struct的key都是大写，map的key可以自定义
	//1.构造struct
	// myJson := new(Json)
	// myJson.Content = "hello"
	// myJson.Title = "world"
	// myJson.User.Name = "Admin"
	// myJson.User.Password = "123"
	//1.构造map
	// myJson := map[string]interface{}{
	// 	"title":   "hello",
	// 	"content": "world",
	// 	"user":    map[string]interface{}{"name": "admin", "password": "123"},
	// }
	myJson := make(map[string]interface{})
	myJson["title"] = "hello"
	myJson["content"] = "world"
	myJson["user"] = map[string]interface{}{
		"name":     "admin",
		"password": 123,
	}

	//2.解析struct到json
	b, err := json.Marshal(myJson)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error no data"))
	}
	//3.返回头
	w.WriteHeader(http.StatusOK)
	//4.返回json
	w.Write(b)

}
