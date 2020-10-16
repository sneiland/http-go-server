package main
 
import (
    "encoding/json"
    "fmt"
    "log"
	"net/http"
	"github.com/gorilla/mux"
)
 
// Article ...
type Article struct {
	ID     string `json:"Id"`
    Title  string `json:"Title"`
    Author string `json:"author"`
    Link   string `json:"link"`
}
 
// Articles ...
var Articles []Article
 
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
 
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}",returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
 
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
 
    for _, article := range Articles {
        if article.ID == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}
 
func main() {
    Articles = []Article{
        Article{Title: "Python Intermediate and Advanced 101",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089KVK23P"},
        Article{Title: "R programming Advanced",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089WH12CR"},
        Article{Title: "R programming Fundamentals",
            Author: "Arkaprabha Majumdar",
            Link:   "https://www.amazon.com/dp/B089S58WWG"},
    }
    handleRequests()
}