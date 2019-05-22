package bots

import (
	"fmt";
	"time";
)

var dictWeekdaysRus = [...]string {
	"воскресенье",
	"понедельник",
	"вторник",
	"среда",
	"четверг",
	"пятница",
	"суббота",
}

var dictMonthesRus = [...]string{
	"dimmy",
	"января",
	"февраля",
	"марта",
	"апреля",
	"мая",
	"июня",
	"июля",
	"августа",
	"сентября",
	"октября",
	"ноября",
	"декабря",
}

type BotRu struct {}

func (bot BotRu) Response(cmd string) (Response string, ExitFlag int) {
	switch cmd {
	case "1","Привет": return bot.SayHello()   , 0
	case "2","Время" : return bot.SayTime()    , 0
	case "3","Дата"  : return bot.SayDate()    , 0
	case "4","День"  : return bot.SayWeekDay() , 0
	case "5","Пока"  : return bot.SayBye()     , 1
	default          : return fmt.Sprintf("Бананов нямя, попробуй что-нибудь еще"), 0
	}
}

func (bot BotRu) SayHello() string {
	return fmt.Sprintf("Привет, я Вася! ;-)")
}

func (bot BotRu) SayTime() string {
	t := time.Now();
	loc, _ := time.LoadLocation("Europe/Minsk")
	t = t.In(loc)
	return fmt.Sprintf("Сейчас %02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

func (bot BotRu) SayDate() string {
	t := time.Now()
	loc, _ := time.LoadLocation("Europe/Minsk")
	t = t.In(loc)
	return fmt.Sprintf("Сегодня %2d %s %d года", t.Day(), dictMonthesRus[t.Month()], t.Year())
}

func (bot BotRu) SayWeekDay() string {
	t := time.Now()
	loc, _ := time.LoadLocation("Europe/Minsk")
	t = t.In(loc)
	return fmt.Sprintf("Сегодня %s", dictWeekdaysRus[t.Weekday()])
}

func (bot BotRu) SayBye() string {
	return fmt.Sprintf("Ты заходи, если чо!")
}
