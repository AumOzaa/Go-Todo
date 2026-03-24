package tools

import (
	"fmt"
)

type returnArticle struct {
	Title       string
	Description string
}

func GetArticle(dateParams string, s string) returnArticle {

	var returnArticle returnArticle
	fmt.Println("In the database")

	returnArticle.Title = "The who got seleceted"
	returnArticle.Description = "The article is coming up soon"

	return returnArticle
}
