package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name  string
	Score int
}

const maxDice = 6
const minDice = 1
const reportString = "%s (%d) rolled a %d, waiting %d sec\n"

func roll() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(maxDice-minDice) + minDice
}

func game(player Player) {
	for {
		roll := roll()
		player.Score += roll

		wait := 6 - roll

		time.Sleep(time.Duration(wait) * time.Second)

		fmt.Printf(reportString, player.Name, player.Score, roll, wait)
	}
}

func main() {
	p1 := Player{Name: "Ed"}
	p2 := Player{Name: "Natasha"}

	go game(p1)
	go game(p2)

	time.Sleep(30 * time.Second)
	fmt.Println("Time up!")
}
