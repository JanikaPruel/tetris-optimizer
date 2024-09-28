package main

import (
	"log"
	"tetris-optimizer/internal/parser"
	"tetris-optimizer/internal/solver"
)

func main() {

	// чтение файла
	content := parser.Readfile("../../data/sample.txt")
	if content == "" {
		log.Fatalf("Failed to read file")
	}

	// парсинг содержимого
	tetrominoes, err := parser.ParseFile(content)
	if err != nil {
		log.Fatalf("Error parsing file: %v", err)
	}

	if len(tetrominoes) == 0 {
		log.Fatalf("No tetrominoes found")
	}
	// Вывод результата парсинга (отладка)
	// for i, tetromino := range tetrominoes {
	// 	fmt.Printf("Tetromino %d:\n", i+1)
	// 	for _, row := range tetromino.Shape {
	// 		fmt.Println(string(row[:]))
	// 	}
	// 	fmt.Println()
	// }

	//
	board := solver.FindMinimalSquare(tetrominoes)
	board.Print()
}
