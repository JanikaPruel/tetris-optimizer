package parser

import (
	"errors"
	"log"
	"os"
	"strings"
)

type Tetromino struct {
	Shape [4][4]rune // Матрица 4х4 для хранения тетромино
}

func Readfile(tetrominos string) string {
	file, err := os.ReadFile(tetrominos)
	if err != nil {
		log.Printf("Error reading file %v", err)
	}

	return string(file)
}

func ParseFile(content string) ([]Tetromino, error) {

	// Разделение содержимого файла на блоки 
	blocks := strings.Split(content, "\n\n")
	tetrominoes := make([]Tetromino, 0, len(blocks))

	for _, block := range blocks {
		block = strings.TrimSpace(block) // удаляем пробелы и переносы строк по краям
		lines := strings.Split(block, "\n") // разделяем блок на строки

		if len(lines) != 4 {
			return nil, errors.New("invalid tetromino format: should be 4 lines")
		}

		var t Tetromino

		for i, line := range lines {
			if len(line) != 4 {
				return nil, errors.New("invalid tetromino format: each line should have 4 characters")
			}
			// преобразуем строку в массив рун и проверяем допустимые символы
			for j, char := range line {
				if char != '#' && char != '.' {
					return nil, errors.New("invalid character tetromino: only '#' and '.' allowed")
				}
				t.Shape[i][j] = char
			}
		}
		// добавляем расписанное тетромино в список
		tetrominoes = append(tetrominoes, t)
	}

	return tetrominoes, nil

	}

