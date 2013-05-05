package govoting

import (
	"appengine"
	"appengine/datastore"

	mux "github.com/gorilla/mux"

	"html/template"
	"net/http"
	"time"
)

type VoteItem struct {
	Submitter      string
	Title          string
	Link           string
	SubmissionTime time.Time
}

// templates variable
var templates = template.Must(template.ParseGlob("templates/*.html"))

func render(res http.ResponseWriter, name string) {
	templates.ExecuteTemplate(res, "header", nil)
	templates.ExecuteTemplate(res, name, nil)
	templates.ExecuteTemplate(res, "footer", nil)

}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	// create context and query the vote items
	context := appengine.NewContext(req)
	query := datastore.NewQuery("VoteItem").Order("-SubmissionTime").Limit(20)

	// create slices to hold the vote items
	vote_items := make([]VoteItem, 0, 20)

	if _, err := query.GetAll(context, &vote_items); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(res, "header", nil)
	templates.ExecuteTemplate(res, "index.html", nil)
	templates.ExecuteTemplate(res, "vote_items.html", vote_items)
	templates.ExecuteTemplate(res, "footer", nil)
}

func submitHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		// create context and prepare vote_item to be saved
		context := appengine.NewContext(req)
		vote_item := VoteItem{
			Submitter:      "Anonymous",
			Title:          req.FormValue("title"),
			Link:           req.FormValue("link"),
			SubmissionTime: time.Now(),
		}

		// save and handle error
		_, err := datastore.Put(context, datastore.NewIncompleteKey(context, "VoteItem", nil), &vote_item)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(res, req, "/", http.StatusFound)
	}

	render(res, "submit.html")
}

func init() {

	//create a new mux router
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/home", rootHandler)
	router.HandleFunc("/submit", submitHandler)

	// register it to the net/http handler
	http.Handle("/", router)
}
