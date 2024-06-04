### Описание задания

В этом проекте вам нужно разработать программу на языке Go, которая принимает путь к текстовому файлу в качестве аргумента. В этом файле содержится список тетромино (фигур из игры Тетрис), которые нужно собрать таким образом, чтобы получился наименьший возможный квадрат. 

### Требования к программе

1. Программа должна успешно компилироваться.
2. Программа должна собирать все тетромино, чтобы создать наименьший квадрат.
3. Программа должна обозначать каждое тетромино в решении заглавными латинскими буквами (A для первого, B для второго и т.д.).
4. В файле должно быть как минимум одно тетромино.
5. В случае некорректного формата тетромино или файла программа должна выводить "ERROR".
6. Программа должна быть написана на языке Go.
7. Код должен соответствовать хорошим практикам.
8. Рекомендуется иметь файлы с тестами для модульного тестирования.

### Пример файла

```text
...#
...#
...#
...#

....
....
..##
..##
```

### Алгоритм выполнения проекта

1. **Создание структуры проекта**
    - Создайте папки и файлы для проекта.
    - Инициализируйте модуль Go.

2. **Чтение файла**
    - Реализуйте функцию для чтения содержимого файла.
    - Обработайте ошибки при чтении файла.

3. **Парсинг тетромино**
    - Реализуйте функцию для парсинга тетромино из текстового файла.
    - Обработайте возможные ошибки формата.

4. **Алгоритм сборки квадрата**
    - Реализуйте алгоритм для нахождения наименьшего квадрата, который может вместить все тетромино.
    - Используйте backtracking или другой подходящий алгоритм.

5. **Вывод результата**
    - Реализуйте функцию для вывода результирующего квадрата с обозначениями тетромино.
    - Обработайте случаи, когда невозможно создать полный квадрат.

6. **Тестирование**
    - Напишите модульные тесты для всех основных функций.
    - Проверьте корректность работы программы с различными входными данными.

### Структура проекта

```
tetris-optimizer/
├── cmd/
│   └── tetris/
│       ├── main.go              # Точка входа в приложение.
├── internal/
│   ├── parser/
│   │   ├── parser.go            # Логика парсинга тетромино из файла.
│   │   └── parser_test.go       # Тесты для парсера.
│   ├── solver/
│   │   ├── solver.go            # Логика для нахождения наименьшего квадрата.
│   │   └── solver_test.go       # Тесты для солвера.
├── data/
│   └── sample.txt               # Пример входного файла.
├── go.mod                       # Модульный файл Go.
├── go.sum                       # Контрольные суммы для зависимостей.
├── Makefile                     # Makefile для сборки и тестирования проекта.
└── README.md                    # Основная документация проекта.
```

### Пример кода

#### main.go

```go
package main

import (
    "fmt"
    "log"
    "os"
    "tetris-optimizer/internal/parser"
    "tetris-optimizer/internal/solver"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . <path_to_file>")
        os.Exit(1)
    }

    path := os.Args[1]
    tetrominoes, err := parser.ParseFile(path)
    if err != nil {
        fmt.Println("ERROR")
        os.Exit(1)
    }

    solution, err := solver.Solve(tetrominoes)
    if err != nil {
        fmt.Println("ERROR")
        os.Exit(1)
    }

    for _, line := range solution {
        fmt.Println(line)
    }
}
```

#### parser.go

```go
package parser

import (
    "bufio"
    "errors"
    "os"
    "strings"
)

func ParseFile(path string) ([]Tetromino, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var tetrominoes []Tetromino
    var currentTetromino []string

    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 {
            if len(currentTetromino) > 0 {
                t, err := NewTetromino(currentTetromino)
                if err != nil {
                    return nil, err
                }
                tetrominoes = append(tetrominoes, t)
                currentTetromino = []string{}
            }
            continue
        }
        currentTetromino = append(currentTetromino, line)
    }

    if len(currentTetromino) > 0 {
        t, err := NewTetromino(currentTetromino)
        if err != nil {
            return nil, err
        }
        tetrominoes = append(tetrominoes, t)
    }

    if len(tetrominoes) == 0 {
        return nil, errors.New("no tetrominoes found")
    }

    return tetrominoes, nil
}

type Tetromino struct {
    Shape []string
}

func NewTetromino(lines []string) (Tetromino, error) {
    if len(lines) != 4 {
        return Tetromino{}, errors.New("invalid tetromino format")
    }

    for _, line := range lines {
        if len(line) != 4 {
            return Tetromino{}, errors.New("invalid tetromino format")
        }
    }

    return Tetromino{Shape: lines}, nil
}
```

#### solver.go

```go
package solver

import (
    "errors"
    "tetris-optimizer/internal/parser"
)

func Solve(tetrominoes []parser.Tetromino) ([]string, error) {
    // Ваш алгоритм для нахождения наименьшего квадрата
    return nil, errors.New("not implemented")
}
```

### Пример Makefile

```Makefile
.PHONY: all build run clean test

all: build

build:
    go build -o bin/tetris-optimizer ./cmd/tetris

run: build
    ./bin/tetris-optimizer data/sample.txt

clean:
    rm -rf bin

test:
    go test ./...
```


