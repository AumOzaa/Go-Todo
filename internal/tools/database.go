package tools

import (
	"fmt"
)

type returnArticle struct {
	Title       string
	Description string
}

func GetArticle(dateParams string, s string) (returnArticle, error) {

	var article returnArticle
	fmt.Println("In the database")

	article.Title = "The who got seleceted"
	article.Description = "The article is coming up soon"

	if article.Title == "" {
		return returnArticle{}, fmt.Errorf("article is empty")
	}

	return article, nil
}
