package main

import (
	"conf/api"
	_ "demo/router"
	"net/http"
)

func main() {
	// setPort()
}
func setPort() {
	println("start server at" + api.AppConfig.String("httpport"))
	//port
	e := http.ListenAndServe(api.AppConfig.String("httpport"), nil)
	if e != nil {
		panic(e)
	}

}
