package bots

type Boter interface {
	Response(cmd string) (Response string, ExitFlag int)
	SayHello() string
	SayTime() string
	SayDate() string
	SayWeekDay() string
	SayBye() string
}

func CreateBot(language string) Boter {
	switch language {
		case "0": return new(BotRu)
		case "1": return new(BotEn)
		default: return nil
	}
}
