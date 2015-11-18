package boards

import (
	// "fmt"
	"github.com/Wmaxlees/go-lg-chess/node"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	matrix "github.com/skelterjohn/go.matrix"
	"image/color"
	"strconv"
)

var singlePawnMove *matrix.DenseMatrix
var singleRookMove *matrix.DenseMatrix
var singleKnightMove *matrix.DenseMatrix
var singleBishopMove *matrix.DenseMatrix
var singleQueenMove *matrix.DenseMatrix
var singleKingMove *matrix.DenseMatrix
var singlePuppyMove *matrix.DenseMatrix

var myPlot *plot.Plot
var total int
var theGoalX int
var theGoalY int

var HoleBoard *matrix.DenseMatrix

func createSubPlots(root *node.Node, pts plotter.XYs, depth int) {
	if pts == nil {
		pts = make(plotter.XYs, depth+1)
	}

	// fmt.Println("Depth: ", depth)
	pts[depth].X = float64(root.GetX())
	pts[depth].Y = float64(root.GetY())

	if depth == 0 {
		if root.GetX() != theGoalX || root.GetY() != theGoalY {
			return
		}

		pts[depth].X = float64(root.GetX())
		pts[depth].Y = float64(root.GetY())

		total++

		plotutil.AddLinePoints(myPlot, pts)
		return
	}

	for _, item := range root.GetChildren() {
		createSubPlots(item, pts, depth-1)
	}
}

func getNextLocations(piece byte, allMoves *matrix.DenseMatrix, ellipse *matrix.DenseMatrix, baseNode *node.Node) []*node.Node {
	result := make([]*node.Node, 0, 20)

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
	case Puppy:
		singleMove = singlePuppyMove
	}
	singleMove = shiftMatrix(singleMove, x-8, y-8)
	singleMove = singleMove.GetMatrix(7, 0, 8, 8)

	// fmt.Println("(", x, ", ", y, ")")
	// fmt.Println("Single Move: \n", singleMove)

	// fmt.Println("\nAll Moves: \n", allMoves)
	// fmt.Println("\nEllipse: ", ellipse)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			// fmt.Println("(", i+1, ", ", 8-j, ") : single: ", singleMove.Get(j, i), "ellipse: ", ellipse.Get(j, i), "all: ", allMoves.Get(j, i))

			if singleMove.Get(j, i) != 0 && ellipse.Get(j, i) != 0 && allMoves.Get(j, i) == float64(next) {
				// fmt.Println("New Child Node: (", i+1, ", ", 8-j, ")")

				var newNode *node.Node
				newNode = new(node.Node)
				newNode.SetX(i + 1)
				newNode.SetY(8 - j)
				newNode.SetStep(next)

				baseNode.AddChild(newNode)
				result = append(result, newNode)
			}
		}
	}

	return result
}

func GetEllipse(piece byte, startX int, startY int, goalX int, goalY int, maxLength int) {
	start := GenerateMoveBoard(piece, startX, startY)
	result := matrix.Sum(start, GenerateMoveBoard(piece, goalX, goalY))

	theGoalX = goalX
	theGoalY = goalY

	// Get lowest #
	min := float64(maxLength)
	// min := float64(100)
	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 8; j++ {
	// 		if result.Get(i, j) < min {
	// 			min = result.Get(i, j)
	// 		}
	// 	}
	// }

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

	// fmt.Println("Ellipse: \n", result)

	nodes := make([][]*node.Node, maxLength+1, maxLength+1)
	nodes[0] = make([]*node.Node, 0, 20)
	nodes[0] = append(nodes[0], root)

	for i := 1; i <= maxLength; i++ {
		// fmt.Println("i: ", i)
		nodes[i] = make([]*node.Node, 0, 20)
		for _, baseNode := range nodes[i-1] {
			// TODO: Improve append to check for duplicate nodes
			nodes[i] = append(nodes[i], getNextLocations(piece, start, result, baseNode)...)
		}
	}

	// fmt.Println(root)

	myPlot, _ = plot.New()
	// Create the board
	createBoard(myPlot)

	createSubPlots(root, nil, maxLength)

	myPlot.X.Min = .5
	myPlot.X.Max = 8.5
	myPlot.Y.Min = .5
	myPlot.Y.Max = 8.5

	// Add the start and end points
	pts := make(plotter.XYs, 2)
	pts[0].X = float64(startX)
	pts[0].Y = float64(startY)
	pts[1].X = float64(goalX)
	pts[1].Y = float64(goalY)
	s, _ := plotter.NewScatter(pts)
	s.GlyphStyle.Color = color.RGBA{R: 0, B: 0, A: 255}
	myPlot.Add(s)

	myPlot.Title.Text = "Total: " + strconv.Itoa(total)

	myPlot.Save(8*vg.Inch, 8*vg.Inch, "results.png")
}

func createBoard(somePlot *plot.Plot) {
	pts := make(plotter.XYs, 35)

	pts[0].X = float64(8.5)
	pts[0].Y = float64(.5)
	pts[1].X = float64(.5)
	pts[1].Y = float64(.5)

	pts[2].X = float64(.5)
	pts[2].Y = float64(1.5)
	pts[3].X = float64(8.5)
	pts[3].Y = float64(1.5)

	pts[4].X = float64(8.5)
	pts[4].Y = float64(2.5)
	pts[5].X = float64(.5)
	pts[5].Y = float64(2.5)

	pts[6].X = float64(.5)
	pts[6].Y = float64(3.5)
	pts[7].X = float64(8.5)
	pts[7].Y = float64(3.5)

	pts[8].X = float64(8.5)
	pts[8].Y = float64(4.5)
	pts[9].X = float64(.5)
	pts[9].Y = float64(4.5)

	pts[10].X = float64(.5)
	pts[10].Y = float64(5.5)
	pts[11].X = float64(8.5)
	pts[11].Y = float64(5.5)

	pts[12].X = float64(8.5)
	pts[12].Y = float64(6.5)
	pts[13].X = float64(.5)
	pts[13].Y = float64(6.5)

	pts[14].X = float64(.5)
	pts[14].Y = float64(7.5)
	pts[15].X = float64(8.5)
	pts[15].Y = float64(7.5)

	pts[16].X = float64(8.5)
	pts[16].Y = float64(8.5)
	pts[17].X = float64(.5)
	pts[17].Y = float64(8.5)

	pts[18].X = float64(.5)
	pts[18].Y = float64(.5)

	pts[19].X = float64(1.5)
	pts[19].Y = float64(.5)
	pts[20].X = float64(1.5)
	pts[20].Y = float64(8.5)

	pts[21].X = float64(2.5)
	pts[21].Y = float64(8.5)
	pts[22].X = float64(2.5)
	pts[22].Y = float64(.5)

	pts[23].X = float64(3.5)
	pts[23].Y = float64(.5)
	pts[24].X = float64(3.5)
	pts[24].Y = float64(8.5)

	pts[25].X = float64(4.5)
	pts[25].Y = float64(8.5)
	pts[26].X = float64(4.5)
	pts[26].Y = float64(.5)

	pts[27].X = float64(5.5)
	pts[27].Y = float64(.5)
	pts[28].X = float64(5.5)
	pts[28].Y = float64(8.5)

	pts[29].X = float64(6.5)
	pts[29].Y = float64(8.5)
	pts[30].X = float64(6.5)
	pts[30].Y = float64(.5)

	pts[31].X = float64(7.5)
	pts[31].Y = float64(.5)
	pts[32].X = float64(7.5)
	pts[32].Y = float64(8.5)

	pts[33].X = float64(8.5)
	pts[33].Y = float64(8.5)
	pts[34].X = float64(8.5)
	pts[34].Y = float64(.5)

	lines, _ := plotter.NewLine(pts)
	lines.LineStyle.Width = vg.Points(1)
	lines.LineStyle.Color = color.RGBA{R: 0, B: 0, A: 255}

	myPlot.Add(lines)
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

	singlePuppyMove = matrix.MakeDenseMatrix([]float64{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, -1, 0, 0, 0, 0, 0, 0, 0,
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
	Puppy  byte = iota
)

func addMovesToBoard(current *matrix.DenseMatrix, newMoves *matrix.DenseMatrix, steps int) *matrix.DenseMatrix {
	// fmt.Println("New Moves: \n", newMoves.String())

	result := current
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			if newMoves.Get(i, j) != 0 && result.Get(i, j) == float64(0) {
				// fmt.Println("    Adding ", steps, " to (", i, ", ", j, ")")
				// if HoleBoard.Get(i, j) > 0 {
				// 	result.Set(i, j, 500)
				// } else {
				result.Set(i, j, float64(steps))
				// }

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
	case Puppy:
		singleMove = singlePuppyMove
	}
	result = shiftMatrix(singleMove, x-8, y-8)

	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			if HoleBoard.Get(j, i) > float64(0) {
				result.Set(j, i, 500)
			}
		}
	}

	// fmt.Println(HoleBoard.String())

	// Get the secondary moves
	for n := 1; n < 20; n++ {
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

	// fmt.Println(result.String())
	result = result.GetMatrix(7, 0, 8, 8)
	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 8; j++ {
	// 		if HoleBoard.Get(i, j) == 1 {
	// 			result.Set(j, i, 500)
	// 		}
	// 	}
	// }

	return result
}
