package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

func GetUrls(url string) (urls []string, err error) {
	file, err := os.Open(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer file.Close()

	err = json.Unmarshal(data, &urls)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for i, url := range urls {
		urls[i] = "https://www.spoonablerecipes.com/common-ingredients-in-" +
			strings.ReplaceAll(url, " ", "-") + "-dishes"
	}

	return urls, nil
}
