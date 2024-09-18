package solver

type Board struct {
	Size int // размер квадрата (доски)
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
		Size: size,
		Cells: cells,
	}

}