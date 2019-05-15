package bots

import (
	"fmt";
	"time";
)

type BotEn struct {}

 func (bot BotEn) Response(cmd string) (Response string, ExitFlag int) {
	switch cmd {
	case "1","Hi"     : return bot.SayHello() , 0
	case "2","Time"   : return bot.SayTime()  , 0
	case "3","Date"   : return bot.SayDate()  , 0
	case "4","WeekDay": return bot.SayWeekDay(), 0
	case "5","Bye"    : return bot.SayBye()    , 1
	default           : return fmt.Sprintf("m-m-m.. I don't know..."), 0
	}
}

 func (bot BotEn) SayHello() string {
	return fmt.Sprintf("Hello, I'm Billy-The-Kid! ;-)")
}

func (bot BotEn) SayTime() string {
	t := time.Now();
	loc, _ := time.LoadLocation("Europe/London")
	t = t.In(loc)
	return fmt.Sprintf("Now is %02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

func (bot BotEn) SayDate() string {
	t := time.Now()
	loc, _ := time.LoadLocation("Europe/London")
	t = t.In(loc)
	return fmt.Sprintf("Today is year %d, %s %2d", t.Year(), t.Month(), t.Day())
}

func (bot BotEn) SayWeekDay() string {
	t := time.Now()
	loc, _ := time.LoadLocation("Europe/London")
	t = t.In(loc)
	return fmt.Sprintf("Today is %s", t.Weekday())
}

func (bot BotEn) SayBye() string {
	return fmt.Sprintf("See you on the other side!")
}