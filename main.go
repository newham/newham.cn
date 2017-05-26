package main

import (
	"net/http"

	"newham.cn/api"

	_ "newham.cn/router"
)

func main() {
	setPort()
	// appengine.Main()
}
func setPort() {
	println("start server at" + api.AppConfig.String("httpport"))
	//port
	e := http.ListenAndServe(api.AppConfig.String("httpport"), nil)
	if e != nil {
		panic(e)
	}

}
