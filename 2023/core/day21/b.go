package day21

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "21b",
		Short: "Day 21, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

var distances = []util.Matrix[int]{}

type Direction int

const (
	Center Direction = iota
	Top
	Bottom
	Left
	Right
	TopLeft
	TopRight
	BottomLeft
	BottomRight
)

func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	tiles := challenge.TileMap()
	startTiles := []util.Point2D{}
	for point, value := range tiles.Iterator() {
		if value == 'S' {
			startTiles = append(startTiles, point) // the start happens to be at the exact centre, i can abuse this fact
			break
		}
	}
	h, w := tiles.GetHeight(), tiles.GetWidth()
	if h != w {
		panic("not implemented for non squares")
	}
	halfH := h/2 + 1

	startTiles = append(startTiles, util.Point2D{Y: 0, X: startTiles[0].X})
	startTiles = append(startTiles, util.Point2D{Y: h - 1, X: startTiles[0].X})
	startTiles = append(startTiles, util.Point2D{Y: startTiles[0].Y, X: 0})
	startTiles = append(startTiles, util.Point2D{Y: startTiles[0].Y, X: w - 1})
	startTiles = append(startTiles, util.Point2D{Y: 0, X: 0})
	startTiles = append(startTiles, util.Point2D{Y: 0, X: w - 1})
	startTiles = append(startTiles, util.Point2D{Y: h - 1, X: 0})
	startTiles = append(startTiles, util.Point2D{Y: h - 1, X: w - 1})

	for i := 0; i < 9; i++ {
		distances = append(distances, getDistances(startTiles[i], tiles)) // the maximum distance from any start is 195 < 2*131
	}

	// mathematics!
	// i've made several assumptions here
	// the shape is square
	// it takes exactly h steps to traverse the grid left-to-right and top-to-bottom
	// there is a gap of .'s between each copy of the grid too (so we know exactly how long it takes to get to a corner)
	// steps is odd (okay that one is just because I don't want to do boolean arithmetic)
	// n is large enough such that the tiles 2 away from the edges are all completely filled (this introduces the corner cases i spent hours forgetting about)
	// probably some other things to be honest, it's not a very good coding question. It's a logic puzzle where you're not told what pieces
	// you're given, or even the fact that you're meant to be looking for extra pieces.
	steps := 26501365
	x := (steps - halfH) / h
	evenDots := 4*(x/2)*(x/2) + 4*(x/2)
	oddDots := (2*x + 2*x*x) - evenDots
	xEven := x%2 == 0
	n := (steps - halfH) % h
	total :=
		getNumReachable(2*h, false, Center) + // okay so i might be assuming steps is odd here, just, pretend I'm not
			getNumReachable(2*h, false, Left)*oddDots + // center tiles (they enter from multiple directions but its all the same in the end)
			getNumReachable(2*h, true, Left)*evenDots +
			getNumReachable(n, !xEven, Left) + // the extreme edges of the edge tiles
			getNumReachable(n, !xEven, Right) +
			getNumReachable(n, !xEven, Top) +
			getNumReachable(n, !xEven, Bottom) +
			getNumReachable(n+halfH-1, xEven, TopLeft)*(x) + // the diagonals of the edge tiles
			getNumReachable(n+halfH-1, xEven, TopRight)*(x) +
			getNumReachable(n+halfH-1, xEven, BottomLeft)*(x) +
			getNumReachable(n+halfH-1, xEven, BottomRight)*(x) +
			getNumReachable(n-halfH, !xEven, TopLeft)*(x+1) + //the small triangles on the outer fringes i forgot about
			getNumReachable(n-halfH, !xEven, TopRight)*(x+1) +
			getNumReachable(n-halfH, !xEven, BottomLeft)*(x+1) +
			getNumReachable(n-halfH, !xEven, BottomRight)*(x+1)

	return total
}

func getNumReachable(n int, even bool, allowedDirections ...Direction) int {
	reachable := make(map[util.Point2D]bool)
	for _, i := range allowedDirections {
		for point, value := range distances[i].Iterator() {
			if even == (value%2 == 0) && value <= n && value != -1 {
				reachable[point] = true
			}
		}
	}

	return len(reachable)

}

// .: completely filled tiles (easy)
// #: mostly filled tiles
// !: very edge tile
//             !
//            !#!
//           !#.#!
//          !#...#!
//         !#.....#!
//        !#.......#!
//         !#.....#!
//          !#...#!
//           !#.#!
//            !#!
//             !
// steps = 66 + x*131 + 131 + n, n < 131 // is that 65 or 66?
// x = (steps-65-131)/131
// n = (steps - 66) % 131
// num ".": 1 + 2x + 2x^2
// num "#": 1+1+1+1 + (x-1)+(x-1)+(x-1)+(x-1)
// num "!": 1+1+1+1 + x+x+x+x
// max distances
// . : inf
// # : 131 + n
// ! : n

// yes but we forgot something important
//             !
//            !#!
//           !#1#!
//          !#121#!
//         !#12121#!
//        !#1211121#!
//         !#12121#!
//          !#121#!
//           !#1#!
//            !#!
//             !
// num ODD dots: 1, 1+4,1+4,1+4+12,1+4+12 = (1 + 2x + 2x^2) - 4*(x/2)^2 - 4*(x/2)
// num EVEN dots: 0, 0, 8, 8, 8+16,8+16  = 4*(x/2)^2 + 4*(x/2)
// # even = ¬x even
// ! even = x even

// new realisation!
// n is 130, screw the # tiles they'll always be full too

// back to the drawing board
//            ?!?
//           ?!2!?
//          ?!212!?
//         ?!21212!?
//        ?!2121212!?
//        !212111212!
//        ?!2121212!?
//         ?!21212!?
//          ?!212!?
//           ?!2!?
//            ?!?
// i forgot an entire class of points! the tiny little triangles that deal with the overflow from the sides of the diagonal "!"
