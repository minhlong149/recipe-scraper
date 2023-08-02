package utils

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Ingredient struct {
	Name  string
	Image string
}

type Recipe struct {
	Name        string
	Ingredients []Ingredient
}

type Response struct {
	Recipe Recipe
	Error  error
}

var ErrNoRecipeFound = errors.New("no recipe found")

func ScrapeRecipe(url string, recipes chan<- Response) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		recipes <- Response{Error: err}
		return
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s\n", res.StatusCode, res.Status)
		recipes <- Response{Error: err}
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		recipes <- Response{Error: err}
		return
	}

	var recipe Recipe

	content := doc.Find("#contents > div.col-md-8.text-content > div")
	recipeName := content.Find("h1").Text()

	if strings.Contains(recipeName, "Recipes Analyzer") || recipeName == "" {
		recipes <- Response{Error: ErrNoRecipeFound}
		return
	}
	recipe.Name = recipeName[len("Common Ingredients In ") : len(recipeName)-len(" Recipes")]

	content.Find("h2").Each(func(i int, s *goquery.Selection) {
		ingredientName := s.Text()
		ingredientImage := s.Next().AttrOr("src", "")

		recipe.Ingredients = append(recipe.Ingredients, Ingredient{
			Name:  ingredientName,
			Image: ingredientImage,
		})
	})

	recipe.Ingredients = recipe.Ingredients[2 : len(recipe.Ingredients)-2]

	recipes <- Response{Recipe: recipe}
}
