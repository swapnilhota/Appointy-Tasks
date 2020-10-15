package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Article is a struct that groups all necessary field into a single unit
type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	SubTitle  string `json:"subTitle"`
	Content   string `json:"content"`
	TimeStamp string `json:"timeStamp"`
}

var articles []Article = []Article{}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/articles", addArticle).Methods("POST")

	router.HandleFunc("/articles", getAllArticles).Methods("GET")

	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")

	http.ListenAndServe(":5000", router)

}

func getArticle(w http.ResponseWriter, r *http.Request) {

	// get the ID of the article from the route paramter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	var flag int
	flag = 1
	var article Article

	for i := 0; i < len(articles); i++ {

		if id == articles[i].ID {

			article = articles[i]
			flag = 0
			break

		}

	}

	if flag == 1 {

		w.WriteHeader(404)
		w.Write([]byte("Article not found"))
		return

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)

}

func getAllArticles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)

}

func addArticle(w http.ResponseWriter, r *http.Request) {

	//get the Article value from the JSON body

	var newArticle Article

	json.NewDecoder(r.Body).Decode(&newArticle)

	articles = append(articles, newArticle)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(articles)

}
