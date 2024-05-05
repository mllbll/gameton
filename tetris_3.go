package main

import "fmt"

type Figure struct {
	Width, Height int
	Shape         [][]int
}

type Field struct {
	Width, Height int
	Grid          [][]int
}

//эта функция располагает во всех возможных местах в пределах поля

//func placeFigures(field Field, figures []Figure) {
//	largest := figures[0]
//	for i := 0; i <= field.Height-largest.Height; i++ {
//		for j := 0; j <= field.Width-largest.Width; j++ {
//			if canPlace(field, largest, i, j) {
//				placeOnField(&field, largest, i, j)
//				break
//			}
//		}
//	}
//
//	for _, fig := range figures[1:] {
//		for i := 0; i <= field.Height-fig.Height; i++ {
//			for j := 0; j <= field.Width-fig.Width; j++ {
//				if canPlace(field, fig, i, j) {
//					placeOnField(&field, fig, i, j)
//					break
//				}
//			}
//		}
//	}
//
//	displayField(field)
//}

//эта располагает в первом возможном
func placeFigures(field Field, figures []Figure) {
	for _, fig := range figures {
		if canPlace(field, fig, 0, 0) {
			placeOnField(&field, fig, 0, 0)
			break // Размещаем только одну фигуру
		}
	}

	displayField(field)
}

func canPlace(field Field, fig Figure, x, y int) bool {
	// Проверяем, что фигура целиком помещается в заданные границы поля
	if x+fig.Height > field.Height || y+fig.Width > field.Width {
		return false
	}

	for i := 0; i < fig.Height; i++ {
		for j := 0; j < fig.Width; j++ {
			// Проверяем, что клетка поля, куда мы пытаемся разместить фигуру, свободна
			if fig.Shape[i][j] == 1 && field.Grid[x+i][y+j] == 1 {
				return false
			}
		}
	}
	return true
}

func placeOnField(field *Field, fig Figure, x, y int) {
	// Размещаем фигуру на поле с учетом её размера
	for i := 0; i < fig.Height; i++ {
		for j := 0; j < fig.Width; j++ {
			if fig.Shape[i][j] == 1 {
				field.Grid[x+i][y+j] = 1
			}
		}
	}
}

func displayField(field Field) {
	for _, row := range field.Grid {
		for _, cell := range row {
			if cell == 1 {
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

	figures := []Figure{
		//{Width: 3, Height: 3, Shape: [][]int{{1, 1, 1}, {1, 0, 0}, {1, 0, 0}}},
		//{Width: 2, Height: 2, Shape: [][]int{{1, 0}, {1, 1}}},
		//{Width: 1, Height: 1, Shape: [][]int{{1}}},
		//{Width: 2, Height: 2, Shape: [][]int{{1, 0}, {0, 1}}},
		//{Width: 2, Height: 2, Shape: [][]int{{1, 0}, {0, 1}}},
		{Width: 2, Height: 1, Shape: [][]int{{1, 0}}},
		//{Width: 1, Height: 1, Shape: [][]int{{1, 1, 0}, {0, 0, 0}}},
		//{Width: 1, Height: 2, Shape: [][]int{{1, 0, 0}, {1, 0, 0}}},
		//{Width: 3, Height: 1, Shape: [][]int{{1, 1, 1}}},
	}

	placeFigures(field, figures)
}
