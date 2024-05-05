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
	for _, fig := range figures {
		placeFigure(&field, fig)
	}

	displayField(field)
}

func main() {
	field := Field{Width: 8, Height: 11}
	field.Grid = make([][]int, field.Height)
	for i := range field.Grid {
		field.Grid[i] = make([]int, field.Width)
	}

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
		{
			Width:  3,
			Height: 3,
			Shape: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
		{
			Width:  3,
			Height: 1,
			Shape: [][]int{
				{1, 1, 1},
			},
		},
		{
			Width:  3,
			Height: 1,
			Shape: [][]int{
				{1, 1, 1},
			},
		},
		{
			Width:  3,
			Height: 1,
			Shape: [][]int{
				{1, 1, 1},
			},
		},
		{
			Width:  5,
			Height: 1,
			Shape: [][]int{
				{1, 1, 1, 1, 1},
			},
		},
		//{
		//	Width:  5,
		//	Height: 1,
		//	Shape: [][]int{
		//		{1, 1, 1, 1, 1},
		//	},
		//},
		{
			Width:  3,
			Height: 3,
			Shape: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
		{
			Width:  1,
			Height: 1,
			Shape: [][]int{
				{1},
			},
		},
		//{
		//	Width:  3,
		//	Height: 4,
		//	Shape: [][]int{
		//		{1, 1, 1},
		//		{1, 0, 1},
		//		{1, 0, 1},
		//		{0, 0, 0},
		//	},
		//},
		//{
		//	Width:  3,
		//	Height: 3,
		//	Shape: [][]int{
		//		{0, 1, 0},
		//		{0, 1, 0},
		//		{1, 1, 1},
		//	},
		//},
	}

	placeFigures(field, figures)
}

func placeFigure(field *Field, fig Figure) {
	// Находим максимальную высоту, на которую можно опустить фигуру
	maxDrop := field.Height - fig.Height

	// Перебираем возможные позиции фигуры
	for drop := 0; drop <= maxDrop; drop++ {
		for x := 0; x <= field.Width-fig.Width; x++ {
			if canPlace(field, fig, x, drop) {
				// Если нашли подходящее место, размещаем фигуру
				placeOnField(field, fig, x, drop)
				return
			}
		}
	}
}

func canPlace(field *Field, fig Figure, x, y int) bool {
	// Проверяем, что все клетки фигуры вписываются на поле, не пересекают другие фигуры и не находятся на уже занятых клетках
	for i := 0; i < fig.Height; i++ {
		for j := 0; j < fig.Width; j++ {
			if fig.Shape[i][j] == 1 {
				// Проверяем, находится ли текущая клетка на поле
				if y+i >= field.Height || x+j >= field.Width {
					return false
				}
				// Проверяем, свободна ли текущая клетка и нет ли на ней другой фигуры
				if field.Grid[y+i][x+j] == 1 {
					return false
				}
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
	fig.X = x
	fig.Y = y
}

func displayField(field Field) {
	for i := range field.Grid {
		for j := range field.Grid[i] {
			if field.Grid[i][j] == 1 {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
