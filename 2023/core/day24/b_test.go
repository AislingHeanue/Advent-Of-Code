package day24

import (
	"testing"

	"github.com/AislingHeanue/Advent-Of-Code/2023/core"
	"github.com/AislingHeanue/Advent-Of-Code/2023/util"
	"github.com/aclements/go-z3/z3"

	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()
	util.ForceNoWindow = true
	input := core.FromLiteral(`19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`)
	// input = core.FromFile()

	result := partB(input)

	require.Equal(t, 47, result)
	util.AwaitClosure()
}

func TestZ3(t *testing.T) {
	config := z3.NewContextConfig()
	ctx := z3.NewContext(config)

	s := z3.NewSolver(ctx)

	px := ctx.Const("px", ctx.IntSort()).(z3.Int)
	py := ctx.Const("py", ctx.IntSort()).(z3.Int)
	s.Assert(px.Add(py).Eq(ctx.FromInt(5, ctx.IntSort()).(z3.Int)))
	s.Assert(px.Sub(py).Eq(ctx.FromInt(5, ctx.IntSort()).(z3.Int)))
}
