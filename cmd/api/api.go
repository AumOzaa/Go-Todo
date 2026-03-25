package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AumOzaa/Go-Todo/internal/tools"
	// "github.com/AumOzaa/Go-Todo/models"
	"github.com/go-chi/chi"

	// "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	r := chi.NewRouter()

	fmt.Println("API Service Started....")

	r.Get("/articles/{date}-{slug}", func(w http.ResponseWriter, r *http.Request) {
		dateParam := chi.URLParam(r, "date")
		slugParam := chi.URLParam(r, "slug")

		article, err := tools.GetArticle(dateParam, slugParam)

		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(fmt.Sprintf("error fetching article %s-%s : %v", dateParam, slugParam, err)))

			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Printf("%v", article)
		json.NewEncoder(w).Encode(article)

	})

	r.Get("/list", func(w http.ResponseWriter, r *http.Request) {
		// var todos []models.Todo

		for i := 0; i < len(tools.MockTodos); i++ {
			fmt.Println(i)
		}

		fmt.Println(tools.MockTodos)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tools.MockTodos)

	})

	http.ListenAndServe(":8000", r)
}
