package main

import (
	"fmt"

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
			fmt.Println(gameStyle.Foreground(lipgloss.Color("#FF0000")).
						Render("\nInvalid choice. Please enter a valid option."))
		}
	}

	fmt.Printf(gameStyle.Bold(true).Render("\nGreat! You have selected the %s difficulty level. You now have %d chances to guess the number.\n"), choices[choice-1], getChances(choice))
	fmt.Println()
	Chances = getChances(choice)
	TargetNumber = GenerateRandomNumber(1, 100)
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
