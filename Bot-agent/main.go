package main

import (
	"fmt"
	"bot-agent/bot"
)

func main() {

	var input string

	print("Выберите язык.   0 - русский\n")
	print("Choose language. 1 - English\n")

	_, err := fmt.Scanln(&input)
	if err != nil {
		print("Error while reading.\n")
		return;
	}

	b := bot.CreateBot(input)
	if b == nil {
		print("Wrong choice. Next time try another\n");
		return;
	}

	b.SayHello()
	ListenAndResponse(b)
}

func ListenAndResponse (b bot.Boter) {
	ExitFlag := 0
	input := ""

	for ; ExitFlag == 0; {
		_, err := fmt.Scanln(&input)
		if err != nil { print("Error while reading.\n"); continue;}

		ExitFlag = b.Response(input)
	}
}