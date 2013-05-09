package govoting

import (
	"appengine"
	"appengine/datastore"

	mux "github.com/gorilla/mux"

	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

const debug = true

type Vote struct {
	Submitter      string
	Title          string
	Description    string
	SubmissionTime time.Time
	ID             int64
}

type VoteItem struct {
	Submitter      string
	Title          string
	Link           string
	SubmissionTime time.Time
	ID             int64
}

type VoteItemComments struct {
	Submitter      string
	Comment        string
	SubmissionTime time.Time
}

// templates variable
var templates = template.Must(template.ParseGlob("templates/*.html"))

func handleError(res http.ResponseWriter, err error, status_code int) {
	if status_code == 0 {
		status_code = http.StatusInternalServerError
	}
	if err != nil {
		http.Error(res, err.Error(), status_code)
	}
}

func renderTemplate(res http.ResponseWriter, name string, i interface{}) {
	err := templates.ExecuteTemplate(res, name, i)
	handleError(res, err, 0)
}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	// create context and query the vote items
	context := appengine.NewContext(req)

	query := datastore.NewQuery("Vote").Order("-SubmissionTime").Limit(20)

	votes := make([]Vote, 0, 20)

	//query all votes
	keys, err := query.GetAll(context, &votes)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	//attach the key queried to the struct
	for index := range votes {
		votes[index].ID = keys[index].IntID()
	}

	renderTemplate(res, "header", nil)
	renderTemplate(res, "index.html", nil)
	renderTemplate(res, "votes.html", votes)
	renderTemplate(res, "footer", nil)
}

func voteCreateHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		context := appengine.NewContext(req)
		vote := Vote{
			Submitter:      "Anonymous",
			Title:          req.FormValue("title"),
			Description:    req.FormValue("description"),
			SubmissionTime: time.Now(),
		}

		_, err := datastore.Put(context, datastore.NewIncompleteKey(context, "Vote", nil), &vote)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(res, req, "/", http.StatusFound)

	}
	renderTemplate(res, "header", nil)
	renderTemplate(res, "new_vote.html", nil)
	renderTemplate(res, "footer", nil)
}

func voteDetailHandler(res http.ResponseWriter, req *http.Request) {
	// create a context
	context := appengine.NewContext(req)

	//get the variable
	urlVar := mux.Vars(req)
	vote_id, _ := strconv.ParseInt(urlVar["vote_id"], 10, 64)

	//construct key and variable to hold vote
	var vote Vote
	key := datastore.NewKey(context, "Vote", "", vote_id, nil)

	// query and attach the key
	datastore.Get(context, key, &vote)
	vote.ID = key.IntID()

	query := datastore.NewQuery("VoteItem").Order("-SubmissionTime").Ancestor(key).Limit(20)
	vote_items := make([]VoteItem, 0, 20)

	//query all votes
	keys, err := query.GetAll(context, &vote_items)
	handleError(res, err, 0)
	for index := range vote_items {
		vote_items[index].ID = keys[index].IntID()
	}

	renderTemplate(res, "header", nil)
	renderTemplate(res, "vote_detail.html", vote)
	renderTemplate(res, "vote_items.html", vote_items)
	renderTemplate(res, "footer", nil)

}

func voteItemCreateHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// create a context
		context := appengine.NewContext(req)

		//get the variable
		urlVar := mux.Vars(req)
		vote_id, _ := strconv.ParseInt(urlVar["vote_id"], 10, 64)
		parent_key := datastore.NewKey(context, "Vote", "", vote_id, nil)

		// create context and prepare vote_item to be saved
		vote_item := VoteItem{
			Submitter:      "Anonymous",
			Title:          req.FormValue("title"),
			Link:           req.FormValue("link"),
			SubmissionTime: time.Now(),
		}

		// save and handle error
		_, err := datastore.Put(context, datastore.NewIncompleteKey(context, "VoteItem", parent_key), &vote_item)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(res, req, "/vote/"+strconv.FormatInt(vote_id, 10), http.StatusFound)
	}
	renderTemplate(res, "header", nil)
	renderTemplate(res, "submit.html", nil)
	renderTemplate(res, "footer", nil)
}

func voteItemHandler(res http.ResponseWriter, req *http.Request) {

	context := appengine.NewContext(req)

	urlVar := mux.Vars(req)
	vote_item_id, _ := strconv.ParseInt(urlVar["vote_item_id"], 10, 64)

	var vote_item VoteItem

	datastore.Get(context, datastore.NewKey(context, "VoteItem", "", vote_item_id, nil), &vote_item)

	if debug == true {
		log.Println(vote_item)
	}

	templates.ExecuteTemplate(res, "header", nil)
	templates.ExecuteTemplate(res, "vote_item_detail.html", vote_item)
	templates.ExecuteTemplate(res, "footer", nil)
}

func init() {

	//create a new mux router
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/vote/create", voteCreateHandler)
	router.HandleFunc("/vote/{vote_id}", voteDetailHandler)

	router.HandleFunc("/vote/{vote_id}/items/create", voteItemCreateHandler)

	// router.HandleFunc("/submit", submitHandler)
	router.HandleFunc("/{vote_item_id}", voteItemHandler)

	// register it to the net/http handler
	http.Handle("/", router)

}
