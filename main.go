package main

import (
	"fmt"
	"strconv"
	"flag"

	"github.com/charmbracelet/lipgloss"
)

var (
	Chances int
	TargetNumber int
	Attempts int
)

var (
	titleStyle   = lipgloss.NewStyle().Italic(true).
					Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4"))
	gameStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))
	ColdStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8C12")) 	// orange
	HotStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFDF12"))  	// yellow
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#adfabb"))  	// green // #70e686
	FailStyle	 = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))  	// red
)

func main() {

	fmt.Println(titleStyle.Render("Welcome to Num-Guesser!"))
	fmt.Println(gameStyle.Underline(true).Render("\nRules: -"))
	fmt.Println(gameStyle.Render("1. I'm thinking of a number between 1 and 100. You can change this range in the settings."))
	fmt.Println(gameStyle.Render("2. You have to guess the number in as few attempts as possible in order to win."))
	fmt.Println(gameStyle.Render("3. After each guess, you'll be told if your guess is too high, too low, or correct."))

	var settings = flag.Bool("custom", false, "Change range settings before starting the game")
	flag.Parse()
	if *settings {
		var min, max int
		fmt.Print(gameStyle.Underline(true).Render("\nEnter minimum number for range: "))
		_, err := fmt.Scanln(&min)
		if err != nil {
			fmt.Println(FailStyle.Render("Invalid input. Using default minimum value of 1."))
			min = 1
		}

		fmt.Print(gameStyle.Underline(true).Render("Enter maximum number for range: "))
		_, err = fmt.Scanln(&max)
		if err != nil || max <= min {
			fmt.Println(FailStyle.Render("Invalid input. Using default maximum value of 100."))
			max = 100
		}

		ChangeRangeSettings(min, max)
		fmt.Printf(SuccessStyle.Render("\nRange updated! Guess a number between %d and %d.\n"), Min, Max)
	}

	choice := 0
	choices := []string{"Easy", "Medium", "Hard"}

	for ;choice < 1 || choice > 3; {
		fmt.Println(gameStyle.Bold(true).Render("\nSelect difficulty level: -"))
		fmt.Println(gameStyle.Render("1. Easy (10 chances)"))
		fmt.Println(gameStyle.Render("2. Medium (5 chances)"))
		fmt.Println(gameStyle.Render("3. Hard (3 chances)"))

		fmt.Print(gameStyle.Underline(true).Render("\nEnter your choice (1-3): "))

		// scanner := bufio.NewScanner(os.Stdin)
		// scanner.Scan()
		// choiceStr := strings.TrimSpace(scanner.Text())
		_, err := fmt.Scanln(&choice)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println(FailStyle.Render("\nInvalid choice. Please enter a valid option."))
		}
	}

	choiceStr := lipgloss.NewStyle().Bold(true).Render(choices[choice-1])
	chancesStr   := lipgloss.NewStyle().Bold(true).Render(strconv.Itoa(getChances(choice)))
	fmt.Printf(gameStyle.Render("\nGreat! You have selected the %s difficulty level. You now have %s chances to guess the number.\n"), choiceStr, chancesStr)
	fmt.Println()
	Chances = getChances(choice)
	TargetNumber = GenerateRandomNumber(Min, Max)
	GameLoop()
}

func getChances(choice int) int {
	switch choice {
	case 1:
		return 10
	case 2:
		return 5
	case 3:
		return 3
	default:
		return 0
	}
}
