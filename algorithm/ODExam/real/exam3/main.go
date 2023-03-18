package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line1 := strings.Split(sc.Text(), " ")
	m, _ := strconv.Atoi(line1[0])
	n, _ := strconv.Atoi(line1[1])

	sc.Scan()
	var inputs []string
	inputs = strings.Split(sc.Text(), " ")

	maze := make([][]int, m)
	for i := 0; i < m; i++ {
		maze[i] = make([]int, 0)
	}

	i := 0
	for j := range inputs {
		if len(maze[i]) == n {
			i++
		}
		num, _ := strconv.Atoi(inputs[j])
		maze[i] = append(maze[i], num)
	}

	sc.Scan()
	line3 := strings.Split(sc.Text(), " ")
	ti, _ := strconv.Atoi(line3[0])
	tj, _ := strconv.Atoi(line3[1])
	fmt.Println(getSignal(maze, m, n, ti, tj))
}

func getSignal(maze [][]int, m, n, ti, tj int) int {
	q := make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if maze[i][j] > 0 {
				q = append(q, [2]int{i, j})
				break
			}
		}
	}

	steps := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		cur := q[0]
		i, j := cur[0], cur[1]
		if maze[i][j] == 1 {
			break
		}
		for _, step := range steps {
			moveX := i + step[0]
			moveY := j + step[1]
			if moveX >= 0 && moveY >= 0 && moveX < m && moveY < n && maze[moveX][moveY] == 0 {
				maze[moveX][moveY] = maze[i][j] - 1
				q = append(q, [2]int{moveX, moveY})
			}
		}
		q = q[1:]
	}
	return maze[ti][tj]
}
