package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/set"
	"github.com/mitchellh/iochan"
)

func main() {
	f, _ := os.Open("./input.txt")
	contraption := parseInput(f)
	illuminatedCellsPart1 := simulateLightBounces(contraption, Ray{Vec2{0, 0}, Vec2{1, 0}})
	illuminatedCellsPart2 := 0

	for entryRay := range allEntryRays(len(contraption[0]), len(contraption)) {
		illuminatedForThisEntry := simulateLightBounces(contraption, entryRay)
		illuminatedCellsPart2 = max(illuminatedCellsPart2, illuminatedForThisEntry)
	}

	fmt.Printf("Number of illuminated cells for part 1: %d\n", illuminatedCellsPart1)
	fmt.Printf("Number of illuminated cells for part 2: %d\n", illuminatedCellsPart2)
}

func allEntryRays(width int, height int) chan Ray {
	result := make(chan Ray)
	go func() {
		// Horizontal edges
		for i := 1; i < width-1; i++ {
			result <- Ray{position: Vec2{i, 0}, direction: Vec2{0, 1}}
			result <- Ray{position: Vec2{i, height - 1}, direction: Vec2{0, -1}}
		}
		// Vertical edges
		for i := 1; i < height-1; i++ {
			result <- Ray{position: Vec2{0, i}, direction: Vec2{1, 0}}
			result <- Ray{position: Vec2{width - 1, i}, direction: Vec2{-1, 0}}
		}
		// Corners
		result <- Ray{position: Vec2{0, 0}, direction: Vec2{1, 0}}
		result <- Ray{position: Vec2{0, 0}, direction: Vec2{0, 1}}
		result <- Ray{position: Vec2{width - 1, 0}, direction: Vec2{-1, 0}}
		result <- Ray{position: Vec2{width - 1, 0}, direction: Vec2{0, 1}}
		result <- Ray{position: Vec2{width - 1, height - 1}, direction: Vec2{-1, 0}}
		result <- Ray{position: Vec2{width - 1, height - 1}, direction: Vec2{0, -1}}
		result <- Ray{position: Vec2{0, height - 1}, direction: Vec2{1, 0}}
		result <- Ray{position: Vec2{0, height - 1}, direction: Vec2{0, -1}}
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
	position  Vec2
	direction Vec2
}

type Vec2 struct {
	x, y int
}

func (v Vec2) Plus(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

func (v Vec2) IsInBoundingBox(width int, height int) bool {
	return v.x >= 0 && v.x < width && v.y >= 0 && v.y < height
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
		runeAtRay := contraption[ray.position.y][ray.position.x]
		switch runeAtRay {
		case '.':
		case '/':
			ray.direction = Vec2{-ray.direction.y, -ray.direction.x}
		case '\\':
			ray.direction = Vec2{ray.direction.y, ray.direction.x}
		case '-':
			if ray.direction.y != 0 {
				ray.direction = Vec2{1, 0}
				rays.Enqueue(Ray{ray.position.Plus(Vec2{-1, 0}), Vec2{-1, 0}})
			}
		case '|':
			if ray.direction.x != 0 {
				ray.direction = Vec2{0, 1}
				rays.Enqueue(Ray{ray.position.Plus(Vec2{0, -1}), Vec2{0, -1}})
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
