package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var articles = []Article{
	Article{
		ID:          "1",
		Title:       "Learning golang",
		Description: "an introduction to the go ecosystem",
		Content:     "AN",
	},

	Article{
		ID:          "2",
		Title:       "Learning GIN",
		Description: "intro to GIN framework",
		Content:     "DA",
	},
}

type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	//Author      *User  `json:"author"`
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, article := range articles {
		if article.ID == params["id"] {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	article := Article{}
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(len(articles) + 1)
	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)
}

func update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	article := Article{}
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = params["id"]

	for i, article := range articles {
		if article.ID == params["id"] {
			articles[i] = article
			break
		}
	}

	json.NewEncoder(w).Encode(article)
}

func delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, article := range articles {
		if article.ID == params["id"] {
			articles = append(articles[:i], articles[:i+1]...)
			break
		}
	}

	json.NewEncoder(w).Encode(articles)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to sumeris 2")
	fmt.Println("Endpoint Hit: home page")
}
func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/articles", getArticles).Methods("GET")
	router.HandleFunc("/api/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/api/articles/", create).Methods("POST")
	router.HandleFunc("/api/articles/{id}", update).Methods("PUT")
	router.HandleFunc("/api/articles/{id}", delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	handleRequest()
}
