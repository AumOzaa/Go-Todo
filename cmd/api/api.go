package main

import (
	// "context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/AumOzaa/Go-Todo/internal/tools"
	"github.com/AumOzaa/Go-Todo/models"

	// "github.com/AumOzaa/Go-Todo/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	// "github.com/go-chi/chi/middleware"
	// "github.com/jackc/pgx"
	// "github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5log"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	r := chi.NewRouter()

	fmt.Println("API Service Started....")

	r.Get("/list", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tools.MockTodos)

	})

	r.Post("/addtodo", func(w http.ResponseWriter, r *http.Request) {
		var newTodo models.Todo

		err := json.NewDecoder(r.Body).Decode(&newTodo)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			render.Status(r, 400)
			render.Render(w, r, models.ErrInvalidRequest())
			return
		}

		fmt.Printf("\n %v", newTodo)

		UpdatedTodo := append(tools.MockTodos, newTodo)

		tools.MockTodos = UpdatedTodo

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tools.MockTodos)
	})

	r.Delete("/remove/{id}", func(w http.ResponseWriter, r *http.Request) {
		// var deleteTodo models.Todo

		todoId, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			fmt.Printf("Error occured %v", err)
			render.Status(r, 400)
			render.Render(w, r, models.ErrInvalidRequest())
			return
		}

		// fmt.Printf("The value to be deleted is %v\n", todoId)

		var finalIndex int
		for i := range tools.MockTodos {
			if todoId == tools.MockTodos[i].Id {
				finalIndex = todoId
				fmt.Printf("The current index is at %v\n", todoId)
			}
		}

		temp := slices.Delete(tools.MockTodos, finalIndex-1, finalIndex)
		fmt.Printf("The current temp is %v\n", temp)
		tools.MockTodos = temp
		fmt.Printf("The current list is %v\n", tools.MockTodos)

	})

	http.ListenAndServe(":8000", r)
}
