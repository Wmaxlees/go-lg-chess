package main

import (
	"bufio"
	"fmt"
	"github.com/Wmaxlees/go-lg-chess/boards"
	matrix "github.com/skelterjohn/go.matrix"
	"os"
	"strconv"
	"strings"
)

var piece byte
var startX int
var startY int
var goalX int
var goalY int
var maxLength int

func main() {
	getOptions()

	boards.InitMatrices()
	boards.GetEllipse(piece, startX, startY, goalX, goalY, maxLength)

	fmt.Println("Output sent to 'results.png'")
}

func getOptions() {
	reader := bufio.NewReader(os.Stdin)

	// Get the piece choice
	for {
		fmt.Println("Choose a Piece:")
		fmt.Println("    Pawn")
		fmt.Println("    Rook")
		fmt.Println("    Knight")
		fmt.Println("    Bishop")
		fmt.Println("    Queen")
		fmt.Println("    King")
		fmt.Println("    Puppy\n")
		fmt.Print(":: ")
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(text)
		text = strings.TrimSpace(text)

		if text == "pawn" {
			piece = boards.Pawn
			break
		} else if text == "rook" {
			piece = boards.Rook
			break
		} else if text == "knight" {
			piece = boards.Knight
			break
		} else if text == "bishop" {
			piece = boards.Bishop
			break
		} else if text == "queen" {
			piece = boards.Queen
			break
		} else if text == "king" {
			piece = boards.King
			break
		} else if text == "puppy" {
			piece = boards.Puppy
			break
		} else {
			fmt.Println("Invalid Piece")
		}
	}

	for {
		fmt.Print("\nStart Position {a-h}{1-8}: ")
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(text)
		text = strings.TrimSpace(text)

		if len(text) > 2 {
			fmt.Println("Invalid Location")
			continue
		}

		if strings.HasPrefix(text, "a") {
			startX = 1
		} else if strings.HasPrefix(text, "b") {
			startX = 2
		} else if strings.HasPrefix(text, "c") {
			startX = 3
		} else if strings.HasPrefix(text, "d") {
			startX = 4
		} else if strings.HasPrefix(text, "e") {
			startX = 5
		} else if strings.HasPrefix(text, "f") {
			startX = 6
		} else if strings.HasPrefix(text, "g") {
			startX = 7
		} else if strings.HasPrefix(text, "h") {
			startX = 8
		} else {
			fmt.Println("Invalid Location")
			continue
		}

		if strings.HasSuffix(text, "1") {
			startY = 1
			break
		} else if strings.HasSuffix(text, "2") {
			startY = 2
			break
		} else if strings.HasSuffix(text, "3") {
			startY = 3
			break
		} else if strings.HasSuffix(text, "4") {
			startY = 4
			break
		} else if strings.HasSuffix(text, "5") {
			startY = 5
			break
		} else if strings.HasSuffix(text, "6") {
			startY = 6
			break
		} else if strings.HasSuffix(text, "7") {
			startY = 7
			break
		} else if strings.HasSuffix(text, "8") {
			startY = 8
			break
		}
		fmt.Println("Invalid Location")
	}

	for {
		fmt.Print("\nGoal Position {a-h}{1-8}: ")
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(text)
		text = strings.TrimSpace(text)

		if len(text) > 2 {
			fmt.Println("Invalid Location")
			continue
		}

		if strings.HasPrefix(text, "a") {
			goalX = 1
		} else if strings.HasPrefix(text, "b") {
			goalX = 2
		} else if strings.HasPrefix(text, "c") {
			goalX = 3
		} else if strings.HasPrefix(text, "d") {
			goalX = 4
		} else if strings.HasPrefix(text, "e") {
			goalX = 5
		} else if strings.HasPrefix(text, "f") {
			goalX = 6
		} else if strings.HasPrefix(text, "g") {
			goalX = 7
		} else if strings.HasPrefix(text, "h") {
			goalX = 8
		} else {
			fmt.Println("Invalid Location")
			continue
		}

		if strings.HasSuffix(text, "1") {
			goalY = 1
			break
		} else if strings.HasSuffix(text, "2") {
			goalY = 2
			break
		} else if strings.HasSuffix(text, "3") {
			goalY = 3
			break
		} else if strings.HasSuffix(text, "4") {
			goalY = 4
			break
		} else if strings.HasSuffix(text, "5") {
			goalY = 5
			break
		} else if strings.HasSuffix(text, "6") {
			goalY = 6
			break
		} else if strings.HasSuffix(text, "7") {
			goalY = 7
			break
		} else if strings.HasSuffix(text, "8") {
			goalY = 8
			break
		}
		fmt.Println("Invalid Location")
	}

	for {
		fmt.Print("\nLength of Trajectories: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		var err error
		maxLength, err = strconv.Atoi(text)
		if err == nil && maxLength > 0 {
			break
		}

		fmt.Println("Length must be a positive integer")
	}

	fmt.Println("\nEnter Holes {a-h}{1-8}")
	fmt.Println("Enter anything else to continue")
	boards.HoleBoard = matrix.MakeDenseMatrix(make([]float64, 225, 225), 15, 15)
	for {
		fmt.Print(":: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if len(text) > 2 {
			break
		}

		var x int
		var y int

		if strings.HasPrefix(text, "a") {
			x = 0
		} else if strings.HasPrefix(text, "b") {
			x = 1
		} else if strings.HasPrefix(text, "c") {
			x = 2
		} else if strings.HasPrefix(text, "d") {
			x = 3
		} else if strings.HasPrefix(text, "e") {
			x = 4
		} else if strings.HasPrefix(text, "f") {
			x = 5
		} else if strings.HasPrefix(text, "g") {
			x = 6
		} else if strings.HasPrefix(text, "h") {
			x = 7
		} else {
			break
		}

		if strings.HasSuffix(text, "1") {
			y = 14
		} else if strings.HasSuffix(text, "2") {
			y = 13
		} else if strings.HasSuffix(text, "3") {
			y = 12
		} else if strings.HasSuffix(text, "4") {
			y = 11
		} else if strings.HasSuffix(text, "5") {
			y = 10
		} else if strings.HasSuffix(text, "6") {
			y = 9
		} else if strings.HasSuffix(text, "7") {
			y = 8
		} else if strings.HasSuffix(text, "8") {
			y = 7
		} else {
			break
		}

		boards.HoleBoard.Set(y, x, 500)
	}
}
