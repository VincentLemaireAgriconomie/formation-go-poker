package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card = rune
type deck struct {
	cards []card
}

type player struct {
	name  string
	token uint
	hand  []card
	bet   uint
}

type board struct {
	deck   deck
	player []player
}

func main() {

	d := deck{
		cards: []card{
		'🂡', '🂢', '🂣', '🂤', '🂥', '🂦', '🂧', '🂨', '🂩', '🂪', '🂫', '🂬', '🂭', '🂮',
		'🂱', '🂲', '🂳', '🂴', '🂵', '🂶', '🂷', '🂸', '🂹', '🂺', '🂻', '🂼', '🂽', '🂾',
		'🃁', '🃂', '🃃', '🃄', '🃅', '🃆', '🃇', '🃈', '🃉', '🃊', '🃋', '🃌', '🃍', '🃎',
		'🃑', '🃒', '🃓', '🃔', '🃕', '🃖', '🃗', '🃘', '🃙', '🃚', '🃛', '🃜', '🃝', '🃞',
		},
	}

	fmt.Printf("%q\n", d.cards[0])

	board := board{
		deck: d,
		player: []player{
			{
				name:  "Alice",
				token: 1000,
				hand:  nil,
				bet:   2,
			},
			{
				name:  "Bob",
				token: 1000,
				hand:  nil,
				bet:   2,
			},
		},
	}

	fmt.Printf("%+v\n", board)
	fmt.Printf("%v\n", string(d.draw()))

}

func (d deck) draw() card {
	rand.Seed(time.Now().Unix())
	return d.cards[rand.Intn(len(d.cards))]
}
