package govoting

import (
	mux "github.com/gorilla/mux"
	"html/template"
	"net/http"
)

// templates variable
var templates = template.Must(template.ParseGlob("templates/*.html"))

func render(res http.ResponseWriter, name string) {
	templates.ExecuteTemplate(res, "header", nil)
	templates.ExecuteTemplate(res, name, nil)
	templates.ExecuteTemplate(res, "footer", nil)

}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	render(res, "index.html")
}

func init() {

	//create a new mux router
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/home", rootHandler)

	// register it to the net/http handler
	http.Handle("/", router)
}
