package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const intro = `WELCOME TO "LUCKYNUMBER"
-FIRST CHOOSE A POSITIVE NUMBER AND PASS IT AS AN ARGUMENT
-THE GAME WILL GENERATE 5 RANDOM NUMBERS, IF YOUR NUMBER MATCHES 1 OF THEM YOU WIN!
----- GOOD LUCK, HAVE FUN -----`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(intro)
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("PLEASE ENTER NUMBERS ONLY")
		return
	}
	if num < 0 {
		fmt.Println("PLEASE ENTER POSITIVE NUMBERS ONLY")
		return
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println("YOUR CHOSEN NUMBER IS: ", num)
	for i := 0; i < 5; i++ {
		rnd := rand.Intn(num + 1)
		fmt.Printf("%d\t", rnd)
		if rnd == num {
			fmt.Println("\nYOU WON!")
			return
		}
	}
	fmt.Println("\nYOU LOST!")
}
