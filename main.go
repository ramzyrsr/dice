package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func dice(sumDice int, sumPlayer int) {
	var pointPlayer []int
	var dicePlayer []int
	var addDice []int

	fmt.Print("Total pemain:", sumPlayer)
	fmt.Println(" dengan masing masing", sumDice, "buah dadu")

	for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
		dicePlayer = append(dicePlayer, sumDice)
		pointPlayer = append(pointPlayer, 0)

		for i := 0; i < sumDice; i++ {
			min, max := 1, 7
			rand.Seed(time.Now().UnixNano())

			number := rand.Intn(max-min) + min

			addDice = append(addDice, number)
		}
	}

	for time := 0; time < 100; time++ {
		var playerWinner int
		var winnerPoint int
		var endGame int

		fmt.Println("Sesi", time+1, ":")
		for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
			if dicePlayer[idPlayer-1] > 0 {
				var allDiceValue string

				fmt.Print("Pemain #", idPlayer)

				for rollDice := 0; rollDice < dicePlayer[idPlayer-1]; rollDice++ {
					diceValue := (rand.Intn(6) + 1)

					if allDiceValue == "" {
						allDiceValue = strconv.Itoa(diceValue)
					} else {
						allDiceValue += ", " + strconv.Itoa(diceValue)
					}

					if diceValue == 6 {
						addDice[idPlayer-1] = addDice[idPlayer-1] - 1
						pointPlayer[idPlayer-1] = pointPlayer[idPlayer-1] + 1
					}

					if diceValue == 1 {
						plusDice := 1
						addDice[idPlayer-1] = addDice[idPlayer-1] - 1
						for i := idPlayer; i < sumPlayer; i++ {
							if dicePlayer[i] > 0 {
								addDice[i] = addDice[i] + plusDice
								plusDice = 0
							}
						}
						for j := 0; j < idPlayer-1; j++ {
							if dicePlayer[j] >= 0 {
								addDice[j] = addDice[j] + plusDice
								plusDice = 0
							}
						}
					}
				}
				fmt.Println(":", allDiceValue)
				continue
			}
		}
		fmt.Print("\nSetelah evaluasi:\n")

		endGame = 0
		for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
			dicePlayer[idPlayer-1] = dicePlayer[idPlayer-1] + addDice[idPlayer-1]
			addDice[idPlayer-1] = 0

			if winnerPoint < pointPlayer[idPlayer-1] {
				playerWinner = idPlayer
				winnerPoint = pointPlayer[idPlayer-1]
			}

			fmt.Print("Pemain #", idPlayer, " memiliki ")
			fmt.Print(dicePlayer[idPlayer-1], " dadu dan ")
			fmt.Print(pointPlayer[idPlayer-1], " point  \n")
			if dicePlayer[idPlayer-1] == 0 {
				endGame = endGame + 1
			}
		}

		if endGame == sumPlayer-1 {
			fmt.Println()
			fmt.Print("Pemenang adalah Pemain #", playerWinner)
			fmt.Println(" dengan", winnerPoint, "point")
			return
		}

		if playerWinner > 0 {
			fmt.Print("Point tertinggi dimiliki Pemain #", playerWinner)
			fmt.Println(" dengan ", winnerPoint, " point")
		}
		fmt.Println()

		if sumPlayer == 1 {
			for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
				fmt.Print("Pemain #", idPlayer, " memiliki ", dicePlayer)
				fmt.Println(" dadu dan ", pointPlayer[idPlayer-1], " point")
			}

			if playerWinner > 0 {
				fmt.Println("Game dimenangkan oleh Pemain #", playerWinner, " dengan ", winnerPoint, " point")
			}
			fmt.Println("Game Over")
			break
		}
	}
}

func main() {
	// set total dice and total player
	dice(4, 3)
}
