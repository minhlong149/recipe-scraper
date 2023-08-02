package utils

import (
	"fmt"
	"log"
	"os"
)

func ExportRecipes(recipes []Recipe) {
	if _, err := os.Stat("recipes.sql"); err == nil {
		err := os.Remove("recipes.txt")
		if err != nil {
			log.Fatalln(err)
		}
	}

	file, err := os.Create("recipes.sql")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	ingredients := make(map[string]string)

	for _, recipe := range recipes {
		for _, ingredient := range recipe.Ingredients {
			ingredientImage, exist := ingredients[ingredient.Name]
			if !exist || ingredientImage == "" {
				ingredients[ingredient.Name] = ingredient.Image
			}
		}
	}

	for name, image := range ingredients {
		_, err := file.WriteString(fmt.Sprintf("INSERT INTO ingredients (name, image)\nVALUES ('%s', '%s');\n\n", name, image))
		if err != nil {
			log.Println(err)
		}
	}

	for _, recipe := range recipes {
		_, err := file.WriteString(fmt.Sprintf("INSERT INTO dishes (name)\nVALUES ('%s');\n\n", recipe.Name))
		if err != nil {
			log.Println(err)
		}

		for _, ingredient := range recipe.Ingredients {
			_, err := file.WriteString(fmt.Sprintf("INSERT INTO recipes (dish_id, ingredient_id)\nVALUES ((SELECT id FROM dishes WHERE name = '%s'), (SELECT id FROM ingredients WHERE name = '%s'));\n\n", recipe.Name, ingredient.Name))
			if err != nil {
				log.Println(err)
			}
		}
	}

}
