package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Min = 1
	Max = 100
)

func ChangeRangeSettings(min, max int) {
	Min = min
	Max = max
}

func GenerateRandomNumber(min, max int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min+1) + min
}

func CheckGuess(guess, target int) (string, bool) {
	if guess < target {
		if target - guess < int((0.01) * (float32(Max-Min))) {
			return HotStyle.Render(fmt.Sprintf("%d is close enough!", guess)), false
		} else if target - guess >= int((0.01) * (float32(Max-Min))) && target - guess < int((0.05) * (float32(Max-Min))) {
			return HotStyle.Render(fmt.Sprintf("Close! %d is slightly lower!", guess)), false
		} else if target - guess >= int((0.1) * (float32(Max-Min))) {
			return ColdStyle.Render(fmt.Sprintf("Way off! %d is too low!", guess)), false
		}
		// return fmt.Sprintf("Incorrect! %d is too low!", guess), false
	} else if guess > target {
		if guess - target < int((0.01) * (float32(Max-Min))) {
			return HotStyle.Render(fmt.Sprintf("%d is close enough!", guess)), false
		} else if guess - target >= int((0.01) * (float32(Max-Min))) && guess - target < int((0.05) * (float32(Max-Min))) {
			return HotStyle.Render(fmt.Sprintf("Close! %d is slightly higher!", guess)), false
		} else if guess - target >= int((0.05) * (float32(Max-Min))) {
			return ColdStyle.Render(fmt.Sprintf("Way off! %d is too high!", guess)), false
		}
		// return fmt.Sprintf("Incorrect! %d is too high!", guess), false
	} else {
		return SuccessStyle.Render(fmt.Sprintf("🎉 Congratulations! You guessed the number in %d attempts!" , Attempts)), true
	}

	return "", false
}

func GameLoop() {
	for ;Chances > 0; Chances-- {
		// fmt.Printf("You have %d chances left. Enter your guess: ", Chances)
		var guess int
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			// fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		result, ok := CheckGuess(guess, TargetNumber)
		Attempts++
		fmt.Println(result)

		if ok {
			// fmt.Println("Congratulations! You've guessed the number!")
			return
		}
	}
	fmt.Println(FailStyle.Render(fmt.Sprintf("😔 Sorry, you've run out of chances. The correct number was %d.\n", TargetNumber)))
}
