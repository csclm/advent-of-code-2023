package day16

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"
	"strings"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/set"
	"github.com/mitchellh/iochan"
)

func Part1(f *os.File) {
	contraption := parseInput(f)
	illuminatedCellsPart1 := simulateLightBounces(contraption, Ray{aoc.NewVec2(0, 0), aoc.NewVec2(1, 0)})
	fmt.Printf("Number of illuminated cells for part 1: %d\n", illuminatedCellsPart1)
}

func Part2(f *os.File) {
	contraption := parseInput(f)
	illuminatedCellsPart2 := 0
	for entryRay := range allEntryRays(len(contraption[0]), len(contraption)) {
		illuminatedForThisEntry := simulateLightBounces(contraption, entryRay)
		illuminatedCellsPart2 = max(illuminatedCellsPart2, illuminatedForThisEntry)
	}
	fmt.Printf("Number of illuminated cells for part 2: %d\n", illuminatedCellsPart2)
}

func allEntryRays(width int, height int) chan Ray {
	result := make(chan Ray)
	go func() {
		// Horizontal edges
		for i := 1; i < width-1; i++ {
			result <- Ray{position: aoc.NewVec2(i, 0), direction: aoc.NewVec2(0, 1)}
			result <- Ray{position: aoc.NewVec2(i, height-1), direction: aoc.NewVec2(0, -1)}
		}
		// Vertical edges
		for i := 1; i < height-1; i++ {
			result <- Ray{position: aoc.NewVec2(0, i), direction: aoc.NewVec2(1, 0)}
			result <- Ray{position: aoc.NewVec2(width-1, i), direction: aoc.NewVec2(-1, 0)}
		}
		// Corners
		result <- Ray{position: aoc.NewVec2(0, 0), direction: aoc.NewVec2(1, 0)}
		result <- Ray{position: aoc.NewVec2(0, 0), direction: aoc.NewVec2(0, 1)}
		result <- Ray{position: aoc.NewVec2(width-1, 0), direction: aoc.NewVec2(-1, 0)}
		result <- Ray{position: aoc.NewVec2(width-1, 0), direction: aoc.NewVec2(0, 1)}
		result <- Ray{position: aoc.NewVec2(width-1, height-1), direction: aoc.NewVec2(-1, 0)}
		result <- Ray{position: aoc.NewVec2(width-1, height-1), direction: aoc.NewVec2(0, -1)}
		result <- Ray{position: aoc.NewVec2(0, height-1), direction: aoc.NewVec2(1, 0)}
		result <- Ray{position: aoc.NewVec2(0, height-1), direction: aoc.NewVec2(0, -1)}
		close(result)
	}()
	return result
}

func parseInput(f *os.File) [][]rune {
	result := make([][]rune, 0)
	for line := range iochan.DelimReader(f, '\n') {
		trimmedLine := strings.TrimSpace(line)
		result = append(result, []rune(trimmedLine))
	}
	return result
}

type Ray struct {
	position  aoc.Vec2
	direction aoc.Vec2
}

func simulateLightBounces(contraption [][]rune, entryRay Ray) int {
	width := len(contraption[0])
	height := len(contraption)
	rays := queue.New()
	evaluated := set.New()
	rays.Enqueue(entryRay) // starting in the top left heading right
	for rays.Len() > 0 {
		ray := rays.Dequeue().(Ray)
		if !ray.position.IsInBoundingBox(width, height) {
			// Ray has left the contraption, so it can't reflect anymore
			continue
		}
		if evaluated.Has(ray) {
			// Already evaluated a different ray that traveled this direction
			// This ray won't illuminate any new cells
			continue
		}
		evaluated.Insert(ray)
		runeAtRay := contraption[ray.position.Y][ray.position.X]
		switch runeAtRay {
		case '.':
		case '/':
			ray.direction = aoc.NewVec2(-ray.direction.Y, -ray.direction.X)
		case '\\':
			ray.direction = aoc.NewVec2(ray.direction.Y, ray.direction.X)
		case '-':
			if ray.direction.Y != 0 {
				ray.direction = aoc.NewVec2(1, 0)
				rays.Enqueue(Ray{ray.position.Plus(aoc.NewVec2(-1, 0)), aoc.NewVec2(-1, 0)})
			}
		case '|':
			if ray.direction.X != 0 {
				ray.direction = aoc.NewVec2(0, 1)
				rays.Enqueue(Ray{ray.position.Plus(aoc.NewVec2(0, -1)), aoc.NewVec2(0, -1)})
			}
		default:
			panic("unrecognized contraption rune")
		}
		ray.position = ray.position.Plus(ray.direction)
		rays.Enqueue(ray)
	}

	illuminated := set.New()
	evaluated.Do(func(eval interface{}) {
		ray := eval.(Ray)
		illuminated.Insert(ray.position)
	})

	return illuminated.Len()
}
