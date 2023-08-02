package main

import (
	"log"
	"net/http"

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

func scrapeRecipe(url string) (recipe Recipe, err error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return Recipe{}, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s\n", res.StatusCode, res.Status)
		return Recipe{}, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return Recipe{}, err
	}

	content := doc.Find("#contents > div.col-md-8.text-content > div")
	recipeName := content.Find("h1").Text()
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

	return recipe, nil
}

func main() {
	phoRecipe, _ := scrapeRecipe("https://www.spoonablerecipes.com/common-ingredients-in-pho-dishes")
	log.Println(phoRecipe)
}
