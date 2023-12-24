package day24

import (
	"fmt"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/aclements/go-z3/z3"
	"github.com/spf13/cobra"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "24b",
		Short: "Day 24, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(core.FromFile()))
		},
	}
}

// prompt: we *really* hope you like linear algebra
func partB(challenge *core.Input) int {
	// uncomment to use util.Matrix.Draw(). util.WindowBeingUsed is a global variable used to tell the code to stop rendering.
	// util.EbitenSetup()
	stones := core.InputMap(challenge, toStone)

	config := z3.NewContextConfig()
	ctx := z3.NewContext(config)

	s := z3.NewSolver(ctx)

	px := ctx.Const("px", ctx.IntSort()).(z3.Int)
	py := ctx.Const("py", ctx.IntSort()).(z3.Int)
	pz := ctx.Const("pz", ctx.IntSort()).(z3.Int)
	vx := ctx.Const("vx", ctx.IntSort()).(z3.Int)
	vy := ctx.Const("vy", ctx.IntSort()).(z3.Int)
	vz := ctx.Const("vz", ctx.IntSort()).(z3.Int)

	for i := range stones[:10] { // checking if 3 stones is enough to fully define the system
		stone := stones[i]
		spx := ctx.FromInt(int64(stone.Position.X), ctx.IntSort()).(z3.Int)
		spy := ctx.FromInt(int64(stone.Position.Y), ctx.IntSort()).(z3.Int)
		spz := ctx.FromInt(int64(stone.Position.Z), ctx.IntSort()).(z3.Int)
		svx := ctx.FromInt(int64(stone.Velocity.X), ctx.IntSort()).(z3.Int)
		svy := ctx.FromInt(int64(stone.Velocity.Y), ctx.IntSort()).(z3.Int)
		svz := ctx.FromInt(int64(stone.Velocity.Z), ctx.IntSort()).(z3.Int)

		s.Assert(spx.Sub(px).Mul(svz.Sub(vz)).Eq(spz.Sub(pz).Mul(svx.Sub(vx))))
		s.Assert(spx.Sub(px).Mul(svy.Sub(vy)).Eq(spy.Sub(py).Mul(svx.Sub(vx))))
	}
	if ok, err := s.Check(); ok {
		x, _, _ := s.Model().Eval(px, true).(z3.Int).AsInt64()
		y, _, _ := s.Model().Eval(py, true).(z3.Int).AsInt64()
		z, _, _ := s.Model().Eval(pz, true).(z3.Int).AsInt64()

		return int(x + y + z)
	} else {
		fmt.Printf("something went seriously wrong: %v\n", err)

		return 0
	}
}
