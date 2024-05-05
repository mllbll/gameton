package main

import (
	"fmt"
)

type Figure struct {
	Width, Height int
	Shape         [][]int
	X, Y          int // Добавляем поля X и Y
}

type Field struct {
	Width, Height int
	Grid          [][]int
}

func placeFigures(field Field, figures []Figure) {
	for i := 0; i < len(figures); i++ {
		fig := figures[i]
		maxContactPoints := 0
		bestX, bestY := 0, 0
		placeInside := false

		for x := 0; x <= field.Height-fig.Height; x++ {
			for y := 0; y <= field.Width-fig.Width; y++ {
				if canPlace(field, fig, x, y) {
					contactPoints := countContactPoints(field, fig, x, y)
					if contactPoints > maxContactPoints {
						maxContactPoints = contactPoints
						bestX, bestY = x, y
						placeInside = false
					}
				}
			}
		}

		// Проверяем, можно ли разместить фигуру внутри других фигур
		if maxContactPoints > 0 {
			for j := 0; j < len(figures); j++ {
				if i != j && canFitInside(figures[j], fig) {
					// Размещаем фигуру внутри уже существующей фигуры
					placeInside = true
					placeOnField(&field, fig, figures[j].X, figures[j].Y)
					break
				}
			}
		}

		if !placeInside && maxContactPoints > 0 {
			// Размещаем фигуру на поле
			placeOnField(&field, fig, bestX, bestY)
			figures[i].X = bestX
			figures[i].Y = bestY
		}
	}

	displayField(field, figures)
}

// Функция для проверки, можно ли разместить фигуру внутри другой фигуры
func canPlaceFigureInside(existingFigure, newFigure Figure) bool {
	// Проверяем, если координаты новой фигуры находятся внутри существующей фигуры
	return newFigure.X >= existingFigure.X &&
		newFigure.Y >= existingFigure.Y &&
		newFigure.X+newFigure.Width <= existingFigure.X+existingFigure.Width &&
		newFigure.Y+newFigure.Height <= existingFigure.Y+existingFigure.Height
}

// Функция для проверки, может ли фигура newFigure поместиться внутри фигуры existingFigure
// Функция для проверки, можно ли разместить одну фигуру внутри другой
func canFitInside(existingFigure, newFigure Figure) bool {
	// Проверяем, если координаты новой фигуры находятся внутри существующей фигуры
	return newFigure.X >= existingFigure.X &&
		newFigure.Y >= existingFigure.Y &&
		newFigure.X+newFigure.Width <= existingFigure.X+existingFigure.Width &&
		newFigure.Y+newFigure.Height <= existingFigure.Y+existingFigure.Height
}

// Функция для подсчета контактных точек с другими фигурами
func countContactPoints(field Field, fig Figure, x, y int) int {
	contactPoints := 0

	// Перебираем клетки вокруг текущей фигуры
	for i := x - 1; i <= x+fig.Height; i++ {
		for j := y - 1; j <= y+fig.Width; j++ {
			// Проверяем, находится ли клетка в пределах поля
			if i >= 0 && i < field.Height && j >= 0 && j < field.Width {
				// Проверяем, не выходит ли индекс за пределы массива fig.Shape
				if i-x >= 0 && i-x < fig.Height && j-y >= 0 && j-y < fig.Width {
					// Если клетка является частью текущей фигуры
					if fig.Shape[i-x][j-y] == 1 {
						contactPoints++
					}
				}
			}
		}
	}

	return contactPoints
}

// Функция для удаления размещенной фигуры с поля
func removeFromField(field *Field, fig Figure, x, y int) {
	for i := 0; i < fig.Height; i++ {
		for j := 0; j < fig.Width; j++ {
			if fig.Shape[i][j] == 1 {
				field.Grid[x+i][y+j] = 0
			}
		}
	}
}

// Функция для вычисления абсолютного значения числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func canPlace(field Field, fig Figure, x, y int) bool {
	if x < 0 || y < 0 || x+fig.Width > field.Width || y+fig.Height > field.Height {
		return false
	}

	for i := 0; i < fig.Height; i++ {
		for j := 0; j < fig.Width; j++ {
			if fig.Shape[i][j] == 1 && field.Grid[y+i][x+j] == 1 {
				return false
			}
		}
	}

	return true
}

func placeOnField(field *Field, fig Figure, x, y int) {
	for i := 0; i < fig.Height; i++ {
		for j := 0; j < fig.Width; j++ {
			if fig.Shape[i][j] == 1 {
				field.Grid[y+i][x+j] = 1
			}
		}
	}
}

func displayField(field Field, figures []Figure) {
	for i := range field.Grid {
		for j := range field.Grid[i] {
			var cellOccupied bool
			for _, fig := range figures {
				if i >= fig.Y && i < fig.Y+fig.Height && j >= fig.X && j < fig.X+fig.Width && fig.Shape[i-fig.Y][j-fig.X] == 1 {
					cellOccupied = true
					break
				}
			}
			if cellOccupied {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func main() {
	field := Field{Width: 8, Height: 11}
	field.Grid = make([][]int, field.Height)
	for i := range field.Grid {
		field.Grid[i] = make([]int, field.Width)
	}

	// Первая фигура: #####
	figures := []Figure{
		{
			Width:  5,
			Height: 5,
			Shape: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0},
			},
		},
		{
			Width:  4,
			Height: 4,
			Shape: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
		},
	}
	//figures := []Figure{
	//	{Width: 4, Height: 4, Shape: [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, X: 0, Y: 0},
	//	{Width: 3, Height: 3, Shape: [][]int{{1, 1, 1}, {1, 0, 0}, {1, 0, 0}}, X: 0, Y: 5},
	//	{Width: 2, Height: 2, Shape: [][]int{{1, 1}, {1, 0}}, X: 0, Y: 9},
	//}

	placeFigures(field, figures)
}
