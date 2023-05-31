package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	players := []int{1, 2, 3, 4}
	for i := 0; i < len(deck); i += 4 {
		player := players[i/4]
		fmt.Printf("Player %d: %d, %d, %d\n", player, deck[i], deck[i+1], deck[i+2])
	}
}
