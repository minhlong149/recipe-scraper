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

	for _, url := range urls {
		recipe, err := utils.ScrapeRecipe(url)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(recipe)
	}
}
