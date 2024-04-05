package main

import (
	"context"
	"fmt"
	"go-translation/app/infrastructure/service"
)

func main() {
	url := "https://deep-translate1.p.rapidapi.com/language/translate/v2"

	c := service.NewClient(url)

	translated, err := c.Translate(context.Background(), "hello", "ja")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(translated)
}
