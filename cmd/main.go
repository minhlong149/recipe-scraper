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

	recipes := make(chan utils.Response)

	for _, url := range urls {
		go utils.ScrapeRecipe(url, recipes)
	}

	for i := 0; i < len(urls); i++ {
		response := <-recipes
		if response.Error != nil {
			log.Println(response.Error)
			continue
		}

		log.Println(response.Recipe)
	}
}
