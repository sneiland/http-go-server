package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Article ...
type Article struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
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
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Println(vars["id"])

	key, err := strconv.Atoi(vars["id"])
	ErrorCheck(err)

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func readDB() {
	db, e := sql.Open("mysql", "webuser:password@/restapi")
	ErrorCheck(e)

	// close database after all work is done
	defer db.Close()

	PingDB(db)

	// query all data
	rows, e := db.Query("select * from articles")
	ErrorCheck(e)

	// declare empty variable
	var article = Article{}

	// iterate over rows
	for rows.Next() {
		e = rows.Scan(&article.Id, &article.Title, &article.Author, &article.Link)
		Articles = append(Articles, article)
		ErrorCheck(e)
	}
}

func main() {
	readDB()

	/*Articles = []Article{
		Article{
			id:     1,
			title:  "Python Intermediate and Advanced 101",
			author: "Arkaprabha Majumdar",
			link:   "https://www.amazon.com/dp/B089KVK23P"},
		Article{
			id:     2,
			title:  "R programming Advanced",
			author: "Arkaprabha Majumdar",
			link:   "https://www.amazon.com/dp/B089WH12CR"},
		Article{
			id:     3,
			title:  "R programming Fundamentals",
			author: "Arkaprabha Majumdar",
			link:   "https://www.amazon.com/dp/B089S58WWG"},
	}

	for _, article := range Articles {
		fmt.Println(article)
	}*/

	handleRequests()
}
