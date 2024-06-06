package main

import (
	"bufio"
	"fmt"
	"os"
	"proj/pkg/match"
	"proj/pkg/player"
	"strconv"
	"strings"
)

// different color schemes for terminal output
const (
	redColor     = "\033[31m"
	greenColor   = "\033[32m"
	yellowColor  = "\033[33m"
	blueColor    = "\033[34m"
	magentaColor = "\033[35m"
	cyanColor    = "\033[36m"
	resetColor   = "\033[0m"
)

// main is the entry point for the Magical Arena application. It presents a console-based menu
// allowing users to enter the arena or exit the application. Once inside the arena, users can
// choose to teleport into matches or exit back to the main menu. The main function utilizes
// a loop to provide a continuous user interface, handling user choices and invoking relevant
// functions for match management.
//
// The color-coded console output enhances the visual experience, and the application logic is
// structured to handle various user inputs and scenarios. The ManageMatchesInArena function is
// responsible for handling the process of entering, conducting, and managing matches within the arena.
func main() {
	for {
		fmt.Println(cyanColor + "Welcome to Magical Arena 1.0!" + resetColor)
		fmt.Println(magentaColor + "Press 1 to enter the arena or press 0 to exit" + resetColor)

		choice, err := getUserInput("Enter your choice: ")
		if err != nil {
			fmt.Println(redColor + "Please enter a valid choice or press 0 to exit" + resetColor)
			continue
		}

		switch choice {
		case 0:
			fmt.Println(redColor + "Exiting the application. Goodbye!" + resetColor)
			return
		case 1:
			fmt.Println(magentaColor + "Entering the arena..." + resetColor)
			fmt.Println(cyanColor + "Welcome to the arena!" + resetColor)
			fmt.Println(yellowColor + "Press 1 to teleport into matches or press 0 to exit" + resetColor)

			// Take user input to enter a match or exit the application
			choice, err = getUserInput("Enter your choice: ")

			if err != nil {
				fmt.Println(redColor + "Error reading user input: " + err.Error() + resetColor)
				continue
			}

			// Entering inside matches
			if choice == 1 {
				// This function will handle the logic of starting matches and concluding them
				ManageMatchesInArena()
			}

			// Handled arena exiting logic
			if choice == 0 {
				fmt.Println(magentaColor + "Exiting the arena." + resetColor)
			} else {
				fmt.Println(redColor + "Invalid choice. Returning to the main menu." + resetColor)
			}
		default:
			fmt.Println(redColor + "Invalid choice. Please enter 0 or 1." + resetColor)
		}
	}
}

// getUserInput prompts the user with the provided message, reads their input
// from the standard input, trims leading/trailing whitespaces, converts the
// input to an integer, and returns the parsed integer choice.
//
// Parameters:
//   - prompt: The message to prompt the user.
//
// Returns:
//   - int: The parsed integer choice.
//   - error: An error, if any.
func getUserInput(prompt string) (int, error) {
	input, err := getStringInput(prompt)
	if err != nil {
		return 0, err
	}

	// Convert the input to an integer
	choice, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %s", input)
	}

	return choice, nil
}

// ManageMatchesInArena initiates the process for entering and conducting matches in the arena.
//
// This function presents the user with options to either enter a new match or exit the arena.
// It prompts the user for input and creates Player instances for both participants.
// The function then validates the attributes of both players and proceeds to create and conduct a new match.
// The match result and round records are stored in a map, and the match number is incremented for each new match.
//
// The function continues running until the user chooses to exit the matches section by entering 0.
//
// Example:
//
//	ManageMatchesInArena()
//
// Note: Ensure that the necessary color constants, getUserInput, getPlayerAttributes, isValidPlayerAttributes,
// and match packages are correctly imported and defined for the proper functioning of this function.
func ManageMatchesInArena() {
	matchRecords := make(map[int]string)
	matchNo := 1

	for {
		fmt.Println(yellowColor + "Press 1 to start a match or press 0 to exit the arena" + resetColor)

		choice, err := getUserInput("Enter your choice: ")
		if err != nil {
			fmt.Println(redColor + "Please enter a valid choice or press 0 to exit" + resetColor)
			return
		}

		switch choice {
		case 0:
			fmt.Println(magentaColor + "Exiting the matches section." + resetColor)
			return
		case 1:
			fmt.Println(cyanColor + "Entering a new match..." + resetColor)

			player1, err := getPlayerAttributes("Player 1")
			if err != nil {
				fmt.Println(redColor + "Error creating Player 1: " + err.Error() + resetColor)
				continue
			}

			player2, err := getPlayerAttributes("Player 2")
			if err != nil {
				fmt.Println(redColor + "Error creating Player 2: " + err.Error() + resetColor)
				continue
			}

			// Validate players attributes
			if !isValidPlayerAttributes(player1, player2) {
				continue
			}

			// Create a new match
			currentMatch := match.NewMatch(player1, player2)

			// Conducting the match
			_, matchResult := match.ConductMatch(currentMatch)

			// Storing the match result and match round records in map
			matchRecords[matchNo] = matchResult

			// Incrementing the match number
			matchNo++

			fmt.Println(greenColor + "Match result: " + matchResult + resetColor)
		default:
			fmt.Println(redColor + "Invalid choice. Please enter 0 or 1." + resetColor)
		}
	}
}

// isValidPlayerAttributes checks if the attributes of two players are within valid ranges to proceed with a match.
// It compares the attack strength of one player against the health of the other player, considering specific conditions.
//
// Parameters:
//   - player1: The first player.
//   - player2: The second player.
//
// Returns:
//   - bool: True if the attributes are within valid ranges, false otherwise.
func isValidPlayerAttributes(player1, player2 *player.Player) bool {
	playerName1, playerHealth1, playerStrength1, playerAttack1 := player.GetPlayerBaseAttributes(player1)
	playerName2, playerHealth2, playerStrength2, playerAttack2 := player.GetPlayerBaseAttributes(player2)

	// Check for unique names of players
	if playerName1 == playerName2 {
		fmt.Println(redColor + "Player names must be unique." + resetColor)
		return false
	}

	// Check for health must be greater than 0
	if playerHealth1 <= 0 || playerHealth2 <= 0 {
		fmt.Println(redColor + "Player health must be greater than 0." + resetColor)
		return false
	}

	// Check for strength must be greater than 0
	if playerStrength1 <= 0 || playerStrength2 <= 0 {
		fmt.Println(redColor + "Player strength must be greater than 0." + resetColor)
		return false
	}

	// Check for attack must be greater than 0
	if playerAttack1 <= 0 || playerAttack2 <= 0 {
		fmt.Println(redColor + "Player attack must be greater than 0." + resetColor)
		return false
	}

	// Check for attack conditions must be following certain conditions
	if playerAttack1*6 <= playerStrength2 {
		fmt.Println(redColor + "Player 1 attack is too low to damage Player 2." + resetColor)
		return false
	}

	if playerAttack2*6 <= playerStrength1 {
		fmt.Println(redColor + "Player 2 attack is too low to damage Player 1." + resetColor)
		return false
	}

	return true
}

// getPlayerAttributes prompts the user to enter attributes for a player and returns a new Player instance.
//
// Parameters:
//   - playerName: The name of the player.
//
// Returns:
//   - *player.Player: A pointer to the newly created Player instance.
//   - error: An error, if any.
func getPlayerAttributes(playerName string) (*player.Player, error) {
	fmt.Printf(cyanColor+"Enter attributes for %s:\n"+resetColor, playerName)

	name, err := getStringInput("Name: ")
	if err != nil {
		return nil, fmt.Errorf("failed to get player name: %w", err)
	}

	health, err := getIntegerInput("Health: ")
	if err != nil {
		return nil, fmt.Errorf("failed to get player health: %w", err)
	}

	strength, err := getIntegerInput("Strength: ")
	if err != nil {
		return nil, fmt.Errorf("failed to get player strength: %w", err)
	}

	attack, err := getIntegerInput("Attack: ")
	if err != nil {
		return nil, fmt.Errorf("failed to get player attack: %w", err)
	}

	return player.NewPlayer(name, health, strength, attack), nil
}

// getIntegerInput prompts the user with the provided message,
// reads their input from the standard input, trims leading/trailing
// whitespaces, and converts the input to an integer.
//
// It relies on getStringInput to obtain user input as a string and
// then attempts to convert it to an integer using strconv.Atoi.
// If the conversion fails, an error is returned.
//
// Parameters:
//   - prompt: The message to prompt the user for input.
//
// Returns:
//   - int: The parsed integer.
//   - error: An error, if any.
func getIntegerInput(prompt string) (int, error) {
	input, err := getStringInput(prompt)
	if err != nil {
		return 0, err
	}

	// Converting the input to an integer
	return strconv.Atoi(input)
}

// ExposeGetIntegerInput is a wrapper function for getIntegerInput to be used in tests
func ExposeGetIntegerInput(prompt string) (int, error) {
	return getIntegerInput(prompt)
}

// getStringInput prompts the user with the provided message,
// reads their input from the standard input, trims leading/trailing
// whitespaces, and returns the resulting string.
//
// Parameters:
//   - prompt: The message to prompt the user for input.
//
// Returns:
//   - string: The user-input string.
//   - error: An error, if any.
func getStringInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Trim leading/trailing whitespaces
	return strings.TrimSpace(input), nil
}

// ExposeGetStringInput is a wrapper function for getStringInput to be used in tests
func ExposeGetStringInput(prompt string) (string, error) {
	return getStringInput(prompt)
}
