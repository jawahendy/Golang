package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Articles []Article

var articles = Articles{
	Article{Title: "judul pertama", Desc: "deskripsi pertama"},
	Article{Title: "judul kedua", Desc: "deskripsike dua"},
}

func main() {

	//middleware

	http.HandleFunc("/", getHome)
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/post-article", withLogging(postArticle))
	http.ListenAndServe(":3001", nil)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testing di home"))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func postArticle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	http.Error(w, "can't Read body", http.StatusInternalServerError)
		// }
		// w.Write([]byte(string(body)))

		var newarticle Article
		err := json.NewDecoder(r.Body).Decode(&newarticle)

		if err != nil {
			fmt.Printf("Ada Error", err)
		}
		articles = append(articles, newarticle)

		json.NewEncoder(w).Encode(articles)

	} else {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Logged koneksi dari", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
