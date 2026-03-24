package main

import (
	"fmt"
	"net/http"

	"github.com/AumOzaa/Go-Todo/internal/tools"
	"github.com/go-chi/chi"
	// "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	r := chi.NewRouter()

	r.Get("/articles/{date}-{slug}", func(w http.ResponseWriter, r *http.Request) {
		dateParam := chi.URLParam(r, "date")
		slugParam := chi.URLParam(r, "slug")

		article, err := tools.GetArticle(dateParam, slugParam)

		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(fmt.Sprintf("error fetching article %s-%s : %v", dateParam, slugParam, err)))

			return
		}

		if article == nil {
			w.WriteHeader(404)
			w.Write([]byte("article not foun"))
			return
		}

		w.Write([]byte(article.Text()))
	})

	http.ListenAndServe(":8000", r)
}
