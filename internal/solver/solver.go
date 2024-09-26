package solver

import (
	"fmt"
	"tetris-optimizer/internal/parser"
)

type Board struct {
	Size  int      // размер квадрата (доски)
	Cells [][]rune // двумерный срез, представляющий клетки доски
}

func NewBoard(size int) *Board {

	// двумермый срез для хранения состояния каждой клетки доски
	cells := make([][]rune, size)

	for i := range cells {
		cells[i] = make([]rune, size)
		// заполняем строку пустыми клетками ('.')
		for j := range cells[i] {
			cells[i][j] = '.' // все клетки изначально пустые
		}
	}

	return &Board{
		Size:  size,
		Cells: cells,
	}

}

// Метод выводит доску на экран
func (b *Board) Print() {

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fmt.Print(string(b.Cells[i][j]))
		}
		fmt.Println()
	}

}

// функция пытается найти минимальный размер квадрата
func FindMinimalSquare(tetrominos []parser.Tetromino) *Board {

	size := calculateMinBoardSize(len(tetrominos))

	for {
		board := NewBoard(size)
		if board.solve(tetrominos, 0) {
			return board
		}
		size++
	}

}

// функция вычисляет минимально возможный размер квадрата
func calculateMinBoardSize(numTetrominos int) int {

	totalCells := numTetrominos * 4

	size := 1

	for size*size < totalCells {
		size++
	}
	return size
}

// функция использует рекурсивный backtracking для размещения всех тетромино
func (b *Board) solve(tetrominos []parser.Tetromino, index int) bool {

	// если все тетромино размещены
	if index == len(tetrominos) {
		return true
	}

	// проходимся по каждой позиции на доске
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.canPlace(tetrominos[index], i, j) { // если можно разместить на текущей позиции
				b.place(tetrominos[index], i, j, rune('A'+index)) // размещаем тетромино

				if b.solve(tetrominos, index+1) { // пробуем разместить след тетромино
					return true
				}
				// убираем тетромино, если не удалось разместить следующее
				b.place(tetrominos[index], i, j, '.')
			}
		}
	}
	// если не удалось разместить текущее тетромино
	return false

}

// размещает тетромино на доске
func (b *Board) place(tetromino parser.Tetromino, x, y int, letter rune) {

	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] != '.' {
				b.Cells[x+i][y+j] = letter
			}
		}
	}

}

// проверяет, можно ли разместить тетромино в указанной позиции
func (b *Board) canPlace(tetromino parser.Tetromino, x, y int) bool {

	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] != '.' {
				if x+i >= b.Size || y+j >= b.Size || b.Cells[x+i][y+j] != '.' {
					return false
				}
			}
		}
	}
	return true
}
