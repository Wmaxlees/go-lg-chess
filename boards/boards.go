package boards

import (
	"fmt"
	"github.com/Wmaxlees/go-lg-chess/node"
	matrix "github.com/skelterjohn/go.matrix"
)

var singlePawnMove *matrix.DenseMatrix
var singleRookMove *matrix.DenseMatrix
var singleKnightMove *matrix.DenseMatrix
var singleBishopMove *matrix.DenseMatrix
var singleQueenMove *matrix.DenseMatrix
var singleKingMove *matrix.DenseMatrix

func getNextLocations(piece byte, allMoves *matrix.DenseMatrix, ellipse *matrix.DenseMatrix, baseNode *node.Node) {
	x := baseNode.GetX()
	y := baseNode.GetY()
	next := baseNode.GetStep() + 1

	// Start with the single move board
	var singleMove *matrix.DenseMatrix
	switch piece {
	case Pawn:
		singleMove = singlePawnMove
	case Rook:
		singleMove = singleRookMove
	case Knight:
		singleMove = singleKnightMove
	case Bishop:
		singleMove = singleBishopMove
	case Queen:
		singleMove = singleQueenMove
	case King:
		singleMove = singleKingMove
	}
	singleMove = shiftMatrix(singleMove, x-8, y-8)
	singleMove = singleMove.GetMatrix(7, 0, 8, 8)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if singleMove.Get(i, j) != 0 && ellipse.Get(i, j) != 0 && allMoves.Get(i, j) == float64(next) {
				var newNode *node.Node
				newNode = new(node.Node)
				newNode.SetX(i)
				newNode.SetY(j)
				newNode.SetStep(next)

				baseNode.AddChild(newNode)
			}
		}
	}

}

func GetEllipse(piece byte, startX int, startY int, goalX int, goalY int) {
	start := GenerateMoveBoard(piece, startX, startY)
	result := matrix.Sum(start, GenerateMoveBoard(piece, goalX, goalY))

	// Get lowest #
	min := float64(100)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if result.Get(i, j) < min {
				min = result.Get(i, j)
			}
		}
	}

	// Remove unnecessary #s
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if result.Get(i, j) > min {
				result.Set(i, j, 0)
			}
		}
	}

	var root *node.Node
	root = new(node.Node)
	root.SetX(startX)
	root.SetY(startY)
	root.SetStep(0)

	getNextLocations(piece, start, result, root)

	fmt.Println(result.String())
}

func InitMatrices() {
	singlePawnMove = matrix.MakeDenseMatrix([]float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, -1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, 15, 15)

	singleRookMove = matrix.MakeDenseMatrix([]float64{
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 1, -1, 1, 1, 1, 1, 1, 1, 1,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
	}, 15, 15)

	singleKnightMove = matrix.MakeDenseMatrix([]float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, -1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, 15, 15)

	singleBishopMove = matrix.MakeDenseMatrix([]float64{
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, -1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	}, 15, 15)

	singleQueenMove = matrix.MakeDenseMatrix([]float64{
		1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1,
		0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
		1, 1, 1, 1, 1, 1, 1, -1, 1, 1, 1, 1, 1, 1, 1,
		0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0,
		1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1,
	}, 15, 15)

	singleKingMove = matrix.MakeDenseMatrix([]float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, -1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, 15, 15)
}

func shiftMatrix(A *matrix.DenseMatrix, x int, y int) *matrix.DenseMatrix {
	result := A

	// fmt.Println("Shift x: ", x)
	// fmt.Println("Shift y: ", y)

	// A*shiftRightOrUp will shift A right
	// shiftRightOrUp*A will shift A up
	shiftRightOrUp := matrix.MakeDenseMatrix([]float64{
		0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, 15, 15)

	// A*shiftLeftOrDown will shift A left
	// shiftLeftOrDown*A will shift A down
	shiftLeftOrDown := matrix.MakeDenseMatrix([]float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
	}, 15, 15)

	// fmt.Println("Start Matrix: ", result.String(), "\n\n")

	if x < 0 {
		for i := x; i < 0; i++ {
			// fmt.Print("<")
			result = matrix.Product(result, shiftLeftOrDown)
			// fmt.Println("Shifted Matrix ", i, ": ", result.String(), "\n\n")
		}
	} else if x > 0 {
		for i := 0; i < x; i++ {
			// fmt.Print(">")
			result = matrix.Product(result, shiftRightOrUp)
			// fmt.Println("Shifted Matrix ", i, ": ", result.String(), "\n\n")
		}
	}
	// fmt.Println()

	if y < 0 {
		for i := y; i < 0; i++ {
			result = matrix.Product(shiftLeftOrDown, result)
			// fmt.Println("Shifted Matrix ", i, ": ", result.String(), "\n\n")
		}
	} else if y > 0 {
		for i := 0; i < y; i++ {
			result = matrix.Product(shiftRightOrUp, result)
			// fmt.Println("Shifted Matrix ", i, ": ", result.String(), "\n\n")
		}
	}

	// fmt.Println("Shifted Matrix: ", result.String())

	return result
}

const (
	Pawn   byte = iota
	Rook   byte = iota
	Knight byte = iota
	Bishop byte = iota
	Queen  byte = iota
	King   byte = iota
)

func addMovesToBoard(current *matrix.DenseMatrix, newMoves *matrix.DenseMatrix, steps int) *matrix.DenseMatrix {
	// fmt.Println("New Moves: \n", newMoves.String())

	result := current
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			if newMoves.Get(i, j) != 0 && result.Get(i, j) == float64(0) {
				// fmt.Println("    Adding ", steps, " to (", i, ", ", j, ")")
				result.Set(i, j, float64(steps))

			}
		}
	}

	return result
}

func GenerateMoveBoard(piece byte, x int, y int) *matrix.DenseMatrix {
	var singleMove *matrix.DenseMatrix
	var result *matrix.DenseMatrix

	// Start with the single move board
	switch piece {
	case Pawn:
		singleMove = singlePawnMove
	case Rook:
		singleMove = singleRookMove
	case Knight:
		singleMove = singleKnightMove
	case Bishop:
		singleMove = singleBishopMove
	case Queen:
		singleMove = singleQueenMove
	case King:
		singleMove = singleKingMove
	}
	result = shiftMatrix(singleMove, x-8, y-8)

	// Get the secondary moves
	for n := 1; n < 8; n++ {
		// fmt.Println("Current State: \n", result.String())
		for i := 0; i < 15; i++ {
			for j := 0; j < 15; j++ {
				// Check if the current position needs to generate it's child moves
				if result.Get(i, j) == float64(n) {
					// Shift the single move matrix
					// fmt.Println("Generating moves from position (", j, ", ", i, ") as ", n)
					result = addMovesToBoard(result, shiftMatrix(singleMove, j-7, 15-(i+8)), n+1)
				}
			}
		}
	}

	// Remove the -1
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			// Check if the current position needs to generate it's child moves
			if result.Get(i, j) == float64(-1) {
				result.Set(i, j, 0)
			}
		}
	}

	// fmt.Println(result.GetMatrix(7, 0, 8, 8).String())
	return result.GetMatrix(7, 0, 8, 8)
}
