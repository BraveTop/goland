package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin(dispatchCoin []string) (int, map[string]int){
	for _, user := range users{
		distribution[user] = 0
		for _, item := range user{
			if string(item) == "e" || string(item) == "E" {
				distribution[user] += 1
				coins -= 1
			}

			if string(item) == "i" || string(item) == "I" {
				distribution[user] += 2
				coins -= 2
			}

			if string(item) == "o" || string(item) == "O" {
				distribution[user] += 3
				coins -= 3
			}

			if string(item) == "u" || string(item) == "U" {
				distribution[user] += 4
				coins -= 4
			}
		}
	}
	return coins, distribution
}

func main() {
	left, value := dispatchCoin(users)
	fmt.Println("剩下：", left)
	fmt.Println("剩下：", value)
}