package main

import (
	"bufio"
	"context"
	"fmt"
	"go-translation/app/infrastructure/service"
	"os"
)

func main() {
	url := "https://deep-translate1.p.rapidapi.com/language/translate/v2"
	var input, target string

	fmt.Print("Please enter the input: ")
	inputScan := bufio.NewScanner(os.Stdin)
	if inputScan.Scan() {
		input = inputScan.Text()
	}

	fmt.Print("Please enter the language to translate to: ")
	targetScan := bufio.NewScanner(os.Stdin)
	if targetScan.Scan() {
		target = targetScan.Text()
	}

	c := service.NewClient(url)

	translated, err := c.Translate(context.Background(), input, target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(translated)
}
