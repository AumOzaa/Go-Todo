package main

import (
	"encoding/json"
	"fmt"
	"github.com/AumOzaa/Go-Todo/internal/tools"
	"github.com/go-chi/chi"
	"net/http"
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

	http.ListenAndServe(":8000", r)
}
