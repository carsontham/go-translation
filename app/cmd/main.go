package main

import (
	"context"
	"fmt"
	"go-translation/app/infrastructure/service"
)

func main() {
	url := "https://deep-translate1.p.rapidapi.com/language/translate/v2"
	var input string
	var target string

	c := service.NewClient(url)
	fmt.Print("Please enter the input: ")
	fmt.Scanln(&input)
	fmt.Print("Please enter the language to translate to: ")
	fmt.Scanln(&target)

	translated, err := c.Translate(context.Background(), input, target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(translated)
}
