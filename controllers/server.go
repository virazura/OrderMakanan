package controllers

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

// RunServer for starting server
func RunServer() {

	// fmt.Println("server started at localhost:8000")
	// http.ListenAndServe(":8000", nil)

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/order-makanan", handlerOrderMakanan)

	var address = "localhost:8000"
	fmt.Printf("Server berjalan di port:%s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func handlerOrderMakanan(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var filepath = path.Join("views", "order-makanan.html")
		var tmpl = template.Must(template.New("result").ParseFiles(filepath))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Batman",
		}

		if err := tmpl.ExecuteTemplate(w, "result", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}
