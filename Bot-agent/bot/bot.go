package bot

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

type Boter interface {
	Response(cmd string) (ExitFlag int)
	SayHello() 
	SayTime()
	SayDate()
	SayWeekDay()
	SayBye()
}

type BotRus struct {}
type BotEn struct {}

func CreateBot(language string) Boter {
	switch language {
		case "0": return new(BotRus)
		case "1": return new(BotEn)
		default: return nil
	}
}

/* *************************************************************
 *			Русский бот
 * ************************************************************* */
func (bot BotRus) Response(cmd string) (ExitFlag int) {
	switch cmd {
	case "1","Привет": bot.SayHello()
	case "2","Время": bot.SayTime()
	case "3","Дата": bot.SayDate()
	case "4","День": bot.SayWeekDay()
	case "5","Пока": bot.SayBye(); return 1
	default: fmt.Println("Бананов нямя, попробуй что-нибудь еще")
	}

	return 0
}

func (bot BotRus) SayHello() {
	fmt.Println("Привет, я Вася! ;-)")
}

func (bot BotRus) SayTime() {
	t := time.Now();
	// Europe/London
	loc, _ := time.LoadLocation("Europe/Minsk")
	t = t.In(loc)
	fmt.Printf("Сейчас %02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())
}

func (bot BotRus) SayDate() {
	t := time.Now()
	loc, _ := time.LoadLocation("Europe/Minsk")
	t = t.In(loc)
	fmt.Printf("Сегодня %2d %s %d года\n", t.Day(), dictMonthesRus[t.Month()], t.Year())
}

func (bot BotRus) SayWeekDay() {
	t := time.Now()
	loc, _ := time.LoadLocation("Europe/Minsk")
	t = t.In(loc)
	fmt.Printf("Сегодня %s\n", dictWeekdaysRus[t.Weekday()])
}

func (bot BotRus) SayBye() {
	fmt.Println("Ты заходи, если чо!")
}

/* *************************************************************
 *			Английский бот
 * ************************************************************* */

 func (bot BotEn) Response(cmd string) (ExitFlag int)  {
	switch cmd {
	case "1","Hi": bot.SayHello()
	case "2","Time": bot.SayTime()
	case "3","Date": bot.SayDate()
	case "4","WeekDay": bot.SayWeekDay()
	case "5","Bye": bot.SayBye(); return 1
	default: fmt.Println("m-m-m.. I don't know...")
	}

	return 0
}

 func (bot BotEn) SayHello() {
	fmt.Println("Hello, I'm Billy-The-Kid! ;-)")
}

func (bot BotEn) SayTime() {
	t := time.Now();
	loc, _ := time.LoadLocation("Europe/London")
	t = t.In(loc)
	fmt.Printf("Now is %02d:%02d:%02d\n", t.Hour(), t.Minute(), t.Second())
}

func (bot BotEn) SayDate() {
	t := time.Now()
	loc, _ := time.LoadLocation("Europe/London")
	t = t.In(loc)
	fmt.Printf("Today is year %d, %s %2d \n", t.Year(), t.Month(), t.Day())
}

func (bot BotEn) SayWeekDay() {
	t := time.Now()
	loc, _ := time.LoadLocation("Europe/London")
	t = t.In(loc)
	fmt.Printf("Today is %s\n", t.Weekday())
}

func (bot BotEn) SayBye() {
	fmt.Println("See you on the other side!")
}