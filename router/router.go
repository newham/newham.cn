package router

import (
	"net/http"

	"newham.cn/controllers"
)

func init() {
	//public
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	//controllers
	http.HandleFunc("/", controllers.IndexController)
	http.HandleFunc("/json", controllers.JsonController)
	http.HandleFunc("/file/add", controllers.FileAddController)
}
