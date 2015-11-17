package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Pawn   byte = iota
	Rook   byte = iota
	Knight byte = iota
	Bishop byte = iota
	Queen  byte = iota
	King   byte = iota
)

var piece byte
var startX byte
var startY byte
var goalX byte
var goalY byte

func main() {
	getOptions()

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
		fmt.Println("    King\n")
		fmt.Print(":: ")
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(text)
		text = strings.TrimSpace(text)

		if text == "pawn" {
			piece = Pawn
			break
		} else if text == "rook" {
			piece = Rook
			break
		} else if text == "knight" {
			piece = Knight
			break
		} else if text == "bishop" {
			piece = Bishop
			break
		} else if text == "queen" {
			piece = Queen
			break
		} else if text == "king" {
			piece = King
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
}
