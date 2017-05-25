package controllers

import "net/http"
import "os"

func FileAddController(w http.ResponseWriter, r *http.Request) {
	// file, err := os.OpenFile("./files/"+r.FormValue("fileName"), os.O_APPEND, 0777)
	file, err := os.Create("./files/" + r.FormValue("fileName"))
	defer file.Close()
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	file.WriteString(r.FormValue("fileContent") + "\n")
	println("add file:" + r.FormValue("fileName"))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.FormValue("fileContent")))
}
