package day23

import (
	"aoc-2023/aoc-lib"
	"os"
)

func parseInput(f *os.File) HikingTrails {
	g, _ := aoc.ReadGrid(f)
	var start aoc.Vec2
	var end aoc.Vec2
	for i := 0; i < g.Width; i++ {
		if g.Get(i, 0) == '.' {
			start = aoc.NewVec2(i, 0)
		}
		if g.Get(i, g.Height-1) == '.' {
			end = aoc.NewVec2(i, g.Height-1)
		}
	}
	return HikingTrails{
		trailMap: g,
		start:    start,
		end:      end,
	}
}
