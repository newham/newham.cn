package main

import (
	"net/http"

	"google.golang.org/appengine"

	"newham-2017/newham.cn/api"

	_ "newham-2017/newham.cn/router"
)

func main() {
	// setPort()
	appengine.Main()
}
func setPort() {
	println("start server at" + api.AppConfig.String("httpport"))
	//port
	e := http.ListenAndServe(api.AppConfig.String("httpport"), nil)
	if e != nil {
		panic(e)
	}

}
