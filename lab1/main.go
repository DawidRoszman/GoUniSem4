package main

import (
	"fmt"
	"math/rand"
)

// (n - 1)/n * 1/(k - 1)
// n - ilosc pudełek
// k - ilosc zamkiętych pudełek

func main() {
	// for i := 1; i < 98; i++ {
	// 	statistics_with_more_boxes(100, i, 10000, true)
	// }
	statistics_with_more_boxes(10, 3, 10000, false)
	statistics_with_more_boxes(10, 7, 10000, true)
}

func statistics_with_more_boxes(num_of_boxes int, boxes_to_eliminate int, num_of_rounds int, strategy bool) {
	var wins int
	var games int

	fmt.Println("Boxes:", num_of_boxes)
	fmt.Println("Boxes to eliminate:", boxes_to_eliminate)
	fmt.Println("Strategy:", strategy)
	fmt.Printf("Playing %d rounds:\n", num_of_rounds)
	for i := 0; i < num_of_rounds; i++ {
		winner_choice := rand.Intn(num_of_boxes) + 1
		user_choice := rand.Intn(num_of_boxes) + 1
		boxes_eliminated := make([]int, boxes_to_eliminate)

		boxes_eliminated = eliminateBoxes(boxes_to_eliminate, boxes_eliminated, num_of_boxes, winner_choice, user_choice)
		user_choice = swap(strategy, num_of_boxes, boxes_eliminated, user_choice)
		if user_choice == winner_choice {
			wins++
		}
		games++

	}

	fmt.Println("Wins:", wins)
	fmt.Println("Loses:", games-wins)
	fmt.Println("Games:", games)
	fmt.Println("Winning percentage:", float64(wins)/float64(games)*100, "%")
	fmt.Println()
}

func swap(strategy bool, num_of_boxes int, boxes_eliminated []int, user_choice int) int {
	if strategy {
		for {

			new_user_choice := rand.Intn(num_of_boxes) + 1
			var found bool

			for _, box := range boxes_eliminated {
				if new_user_choice == box {
					found = true
				}
			}
			if new_user_choice == user_choice {
				found = true
			}
			if !found {
				user_choice = new_user_choice
				break
			}

		}
	}
	return user_choice
}

func eliminateBoxes(boxes_to_eliminate int, boxes_eliminated []int, num_of_boxes int, winner_choice int, user_choice int) []int {
	for j := 0; j < boxes_to_eliminate; j++ {
		for {
			boxes_eliminated[j] = rand.Intn(num_of_boxes) + 1
			if boxes_eliminated[j] != winner_choice && boxes_eliminated[j] != user_choice {
				var found bool
				for k := 0; k < j; k++ {
					if boxes_eliminated[j] == boxes_eliminated[k] {
						found = true
					}
				}
				if !found {
					break
				}

			}
		}
	}
	return boxes_eliminated
}

// func statistics(num_of_rounds int, strategy string) {
// 	var wins int
// 	var games int
//
// 	fmt.Println("Strategy:", strategy)
// 	fmt.Printf("Playing %d rounds:\n", num_of_rounds)
// 	for i := 0; i < num_of_rounds; i++ {
// 		winner_choice := rand.Intn(3) + 1
// 		user_choice := rand.Intn(3) + 1
// 		var eliminated_box int
// 		for {
// 			eliminated_box = rand.Intn(3) + 1
// 			if eliminated_box != winner_choice || eliminated_box != user_choice {
// 				break
// 			}
// 		}
// 		if strategy == "switch" {
// 			user_choice = 6 - user_choice - eliminated_box
// 		}
// 		if user_choice == winner_choice {
// 			wins++
// 		}
// 		games++
//
// 	}
//
// 	fmt.Println("Wins:", wins)
// 	fmt.Println("Loses:", games-wins)
// 	fmt.Println("Games:", games)
// 	fmt.Println("Winning percentage:", float64(wins)/float64(games)*100, "%")
// 	fmt.Println()
// }

func user_game() {
	winner_choice := rand.Intn(3) + 1

	fmt.Println("You have 3 boxes in front of you. Choose one: 1, 2, or 3")
	user_choice := make_choice()

	// Eliminate one box that is not a winner and not the user's choice
	var eliminated_box int
	for {
		eliminated_box = rand.Intn(3) + 1
		if eliminated_box != winner_choice && eliminated_box != user_choice {
			break
		}
	}

	fmt.Println("We have eliminated box", eliminated_box)
	fmt.Println("Would you like to switch your choice? (yes/no)")
	var switch_choice string
	fmt.Scan(&switch_choice)
	if switch_choice == "yes" {
		user_choice = 6 - user_choice - eliminated_box
	}

	if user_choice == winner_choice {
		fmt.Println("Congratulations! You have won!")
	} else {
		fmt.Println("Sorry, you have lost. The winning box was", winner_choice)
	}
}

func make_choice() int {
	var user_choice int

	fmt.Scan(&user_choice)

	return user_choice
}
