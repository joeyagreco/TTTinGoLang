// Tic Tac Toe in GoLang
// Version 2.0
// Joey Greco

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	game := newTicTacToe()
	game.play()

}

type TicTacToe struct {
	// Tic Tac Toe object
	// TicTacToe board as array
	board [9]string
	// Whos turn it is (will be "x" or "o")
	player string
}

func newTicTacToe() *TicTacToe {
	// Creates a new TicTacToe object and returns a pointer to the object
	// Set default starting player to x
	ttt := TicTacToe{player: "x"}
	// Fill the game board with " " to start
	for i := 0; i < 9; i++ {
		ttt.board[i] = " "
	}
	return &ttt
}

func (t TicTacToe) play() {
	// This function is the game loop for TicTacToe
	fmt.Println("Welcome to Tic Tac Toe!")
	// display board
	t.printBoard()
	// infinite loop that only breaks once the game is over
	for {
		// Get the placement from the user
		choice := t.getUserPlacementPosition()
		// Set the piece and update the board
		t.setPiece(choice)
		// display board
		t.printBoard()
		// Check if the game is over and break if so
		if t.gameOver() == 1 {
			// Someone has won
			fmt.Println(t.player + " has won!")
			break
		} else if t.gameOver() == 2 {
			// The game ended in a tie
			fmt.Println("The game ended in a tie")
			break
		}
		// The game isn't over yet, so switch players and continue
		t.switchPlayer()
	}
}

func (t TicTacToe) printBoard() {
	// This prints the t.board as a TicTacToe board in console
	fmt.Println(t.board[0] + " | " + t.board[1] + " | " + t.board[2])
	fmt.Println("_________")
	fmt.Println(t.board[3] + " | " + t.board[4] + " | " + t.board[5])
	fmt.Println("_________")
	fmt.Println(t.board[6] + " | " + t.board[7] + " | " + t.board[8])
}

func (t TicTacToe) getBoard() [9]string {
	// Board getter method
	return t.board
}

func (t TicTacToe) getUserPlacementPosition() int {
	// This prompts the user for a space (1-9) to place their piece on the board
	// It returns an int (1-9)
	// If the user chooses a space that is not empty OR a space that is out of range, we re-prompt
	stdIn := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Where would you like to place your piece?(1-9) : ")
		stdIn.Scan()
		// Convert user input to int
		choice, _ := strconv.Atoi(stdIn.Text())
		// Check if the space is in range
		if choice < 10 && choice > 0 {
			// Check if the space is empty
			if t.board[choice-1] == " " {
				return choice
			}
		}
		// If we reach this, there was invalid input by the user
		fmt.Println("Invalid Input. Must be (1-9) and space must be empty.")
		// Display the board every time we have to ask for a placement
		t.printBoard()

	}
}

func (t *TicTacToe) setPiece(space int) {
	// This sets the current t.player ("x" or "o") on the given space (1-9)
	t.board[space-1] = t.player
}

func (t *TicTacToe) switchPlayer() {
	// This switches the current player
	// "x" -> "o"
	//or
	//"o" -> "x"
	if t.player == "x" {
		t.player = "o"
	} else {
		t.player = "x"
	}
}

func (t TicTacToe) gameOver() int {
	// This returns:
	// 0 if the game isnt over
	// 1 if the game has been won
	// 2 if the game has been tied

	// Check if the game has been won by either player
	if t.board[0] != " " && t.board[0] == t.board[1] && t.board[1] == t.board[2] {
		// Check for top row match
		return 1
	} else if t.board[3] != " " && t.board[3] == t.board[4] && t.board[4] == t.board[5] {
		// Check for middle row match
		return 1
	} else if t.board[6] != " " && t.board[6] == t.board[7] && t.board[7] == t.board[8] {
		// Check for bottom row match
		return 1
	} else if t.board[0] != " " && t.board[0] == t.board[3] && t.board[3] == t.board[6] {
		// Check for left column match
		return 1
	} else if t.board[1] != " " && t.board[1] == t.board[4] && t.board[4] == t.board[7] {
		// Check for middle column match
		return 1
	} else if t.board[2] != " " && t.board[2] == t.board[5] && t.board[5] == t.board[8] {
		// Check for right column match
		return 1
	} else if t.board[0] != " " && t.board[0] == t.board[4] && t.board[4] == t.board[8] {
		// Check for top left to bottom right diagonal match
		return 1
	} else if t.board[2] != " " && t.board[2] == t.board[4] && t.board[4] == t.board[6] {
		// Check for top right to bottom left diagonal match
		return 1
	} else if t.board[0] != " " && t.board[1] != " " && t.board[2] != " " && t.board[3] != " " && t.board[4] != " " && t.board[5] != " " && t.board[6] != " " && t.board[7] != " " && t.board[8] != " " {
		// Check for a tie
		return 2
	}
	// The game isnt over yet
	return 0
}
