package main

import "fmt"

// КРУТИМ ПРОТИВ ЧАСОВОЙ СТРЕЛКИ
type Point struct {
	X int
	Y int
}

func rotate90(points []Point) []Point {
	rotatedPoints := make([]Point, len(points))
	for i, p := range points {
		rotatedPoints[i] = Point{X: p.Y, Y: -p.X}
	}
	return rotatedPoints
}

func rotate270(points []Point) []Point {
	rotatedPoints := make([]Point, len(points))
	for i, p := range points {
		rotatedPoints[i] = Point{X: -p.Y, Y: p.X}
	}
	return rotatedPoints
}

func rotate180(points []Point) []Point {
	rotatedPoints := make([]Point, len(points))
	for i, p := range points {
		rotatedPoints[i] = Point{-p.X, -p.Y}
	}
	return rotatedPoints
}

func display(points []Point) {
	minX, maxX := points[0].X, points[0].X
	minY, maxY := points[0].Y, points[0].Y
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	grid := make([][]bool, maxY-minY+1)
	for i := range grid {
		grid[i] = make([]bool, maxX-minX+1)
	}

	for _, p := range points {
		grid[p.Y-minY][p.X-minX] = true
	}

	for _, row := range grid {
		for _, val := range row {
			if val {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	//точки задаются как в примере в файле геймтона
	points := []Point{{0, 0}, {0, 1}, {1, 0}, {2, 0}}

	fmt.Println("Исходная фигура:")
	display(points)

	fmt.Println("\nПоворот на 90 градусов:")
	rotated90 := rotate90(points)
	display(rotated90)

	fmt.Println("\nПоворот на 180 градусов:")
	rotated180 := rotate180(points)
	display(rotated180)

	fmt.Println("\nПоворот на 270 градусов:")
	rotated270 := rotate270(points)
	display(rotated270)
}
