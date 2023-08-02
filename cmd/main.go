package main

import (
	"log"

	"recipe-scraper/pkg/utils"
)

func main() {
	urls, err := utils.GetUrls("assets/ingredient_data.json")
	if err != nil {
		log.Println(err)
		return
	}

	responses := make(chan utils.Response)

	for _, url := range urls {
		go utils.ScrapeRecipe(url, responses)
	}

	var recipes []utils.Recipe

	for i := 0; i < len(urls); i++ {
		response := <-responses
		if response.Error == nil {
			recipes = append(recipes, response.Recipe)
		}
	}

	utils.ExportRecipes(recipes)
}
