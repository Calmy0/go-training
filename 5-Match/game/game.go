package game

type MatchT struct {
	player1 PlayerT
	player2 PlayerT
}

func (m *MatchT) Start() {
	m.player1 = PlayerT{"Вася", 80}
	m.player2 = PlayerT{"Петя", 50}
	chanBall := make(chan int)
	chanLooser := make(chan string)
	go m.player1.Hit(chanBall, chanLooser)
	go m.player2.Hit(chanBall, chanLooser)

	println("Мяч в игре!")
	chanBall <- 1

	println("Проиграл: ", <-chanLooser)
}
