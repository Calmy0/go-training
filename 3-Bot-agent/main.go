package main

import (
	"fmt"
	"bot-agent/bots"
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

	b := bots.CreateBot(input)
	if b == nil {
		print("Wrong choice. Next time try another\n");
		return;
	}

	fmt.Println(b.SayHello())
	ListenAndResponse(b)
}

func ListenAndResponse (b bots.Boter) {
	ExitFlag := 0
	response :=""
	input := ""

	for ; ExitFlag == 0; {
		_, err := fmt.Scanln(&input)
		if err != nil { print("Error while reading.\n"); continue;}

		response, ExitFlag = b.Response(input)
		fmt.Println(response)
	}
}