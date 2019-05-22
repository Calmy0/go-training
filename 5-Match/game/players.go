package game

import (
	"math/rand"
	"time"
)

type PlayerT struct {
	name  string
	skill int // [0..100]
}

var Random *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	Random = rand.New(s)
}

func (p *PlayerT) Hit(chanBall chan int, chanLooser chan string) {
	var counter int = 0
	for range chanBall { // получает данные пока канал не закрыт
		ballSkill := Random.Intn(100)

		counter++
		if p.skill > ballSkill { // если отбил
			println(counter, " ", p.name, " со скилом ", p.skill, " отбил мяч сложности ", ballSkill)
			chanBall <- 1 // отправляем мяч назад
		} else { // если не отбил
			println(p.name, " со скилом ", p.skill, " НЕ отбил мяч сложности ", ballSkill)
			close(chanBall)      // уведомляем второго
			chanLooser <- p.name // сигнализируем, что пропустил
		}
	}

	// канал закрылся, значит другой пропустил
	return
}
