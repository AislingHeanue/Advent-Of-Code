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

	for i := 0; i < 5; i++ {
		distances = append(distances, getDistances(startTiles[i], tiles)) // the maximum distance from any start is 195 < 2*131
	}

	// mathematics!
	steps := 26501365
	// steps := 9
	x := (steps - halfH - h) / h
	evenDots := 4*(x/2)*(x/2) + 4*(x/2)
	oddDots := (2*x + 2*x*x) - evenDots
	xEven := x%2 == 0
	n := (steps - halfH) % h
	// distancesLocal := distances
	// fmt.Print(distancesLocal)
	total :=
		getNumReachable(2*h, false, Center) + // okay so i might be assuming steps is odd here, just, pretend I'm not
			getNumReachable(2*h, false, Left)*oddDots + // center tiles (they enter from multiple directions but its all the same in the end)
			getNumReachable(2*h, true, Left)*evenDots +
			getNumReachable(n+h, !xEven, Left) + // the extreme edges of the near edge tiles
			getNumReachable(n+h, !xEven, Right) +
			getNumReachable(n+h, !xEven, Top) +
			getNumReachable(n+h, !xEven, Bottom) +
			getNumReachable(n, xEven, Left) + // the extreme edges of the edge tiles
			getNumReachable(n, xEven, Right) +
			getNumReachable(n, xEven, Top) +
			getNumReachable(n, xEven, Bottom) +
			getNumReachable(n+h, !xEven, Left, Top)*(x) + // diagonals of the near edge tiles
			getNumReachable(n+h, !xEven, Right, Top)*(x) +
			getNumReachable(n+h, !xEven, Left, Bottom)*(x) +
			getNumReachable(n+h, !xEven, Right, Bottom)*(x) +
			getNumReachable(n, xEven, Left, Top)*(x+1) + // the diagonals of the edge tiles
			getNumReachable(n, xEven, Right, Top)*(x+1) +
			getNumReachable(n, xEven, Left, Bottom)*(x+1) +
			getNumReachable(n, xEven, Right, Bottom)*(x+1)

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
	// fmt.Println(len(reachable))

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
