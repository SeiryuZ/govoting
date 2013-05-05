package govoting

import (
	"fmt"
	mux "github.com/gorilla/mux"
	// "html/template"
	"net/http"
)

// // templates variable
// var templates = template.Must(template.ParseFiles("templates/index.html"))

// // render functions
// func renderTemplate(res http.ResponseWriter, template string) {
// 	err := templates.ExecuteTemplate(res, template, nil)
// 	if err != nil {
// 		http.Error(res, err.Error(), http.StatusInternalServerError)
// 	}
// }

func rootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, world!")
}

func init() {

	//create a new mux router
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/home", rootHandler)

	// register it to the net http handler
	http.Handle("/", r)
}
