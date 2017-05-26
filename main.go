package main

import (
	"conf/api"
	_ "demo/router"
	"net/http"

	"google.golang.org/appengine"
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
