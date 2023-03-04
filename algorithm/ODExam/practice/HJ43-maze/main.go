package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	m, n := 0, 0
	fmt.Scanln(&m, &n)

	maze := make([][]int, m)
	for i := 0; i < m; i++ {
		maze[i] = make([]int, n)
	}
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < m; i++ {
		var line []string
		for sc.Scan() {
			line = strings.Split(sc.Text(), " ")
			break
		}
		for j, v := range line {
			num, _ := strconv.Atoi(v)
			maze[i][j] = num
		}
	}
	road := findRoad(maze, m, n)
	for i := 1; i < len(road); i++ {
		fmt.Printf("(%c,%c)\n", road[i][0], road[i][2])
	}
}

func findRoad(maze [][]int, m, n int) []string {
	path := []string{"0,0"}
	visited := map[string]bool{"0,0": true}
	res := []string{"0,0"}
	endPoint := fmt.Sprintf("%d,%d", m-1, n-1)
	var dfs func()
	dfs = func() {
		cur := path[len(path)-1]
		x, y := int(cur[0]-'0'), int(cur[2]-'0')
		if maze[x][y] == 1 {
			return
		}
		if cur == endPoint {
			res = append(res, path...)
			return
		}

		if moved := move(cur, 1, m, n); moved != "" && !visited[moved] {
			path = append(path, moved)
			visited[moved] = true
			dfs()
			path = path[:len(path)-1]
			visited[moved] = false
		}

		if moved := move(cur, 2, m, n); moved != "" && !visited[moved] {
			path = append(path, moved)
			visited[moved] = true
			dfs()
			path = path[:len(path)-1]
			visited[moved] = false
		}

		if moved := move(cur, 3, m, n); moved != "" && !visited[moved] {
			path = append(path, moved)
			visited[moved] = true
			dfs()
			path = path[:len(path)-1]
			visited[moved] = false
		}

		if moved := move(cur, 4, m, n); moved != "" && !visited[moved] {
			path = append(path, moved)
			visited[moved] = true
			dfs()
			path = path[:len(path)-1]
			visited[moved] = false
		}
	}
	dfs()
	return res
}

// 1、2、3、4分别是上下左右移动
func move(s string, direction, m, n int) string {
	var res string
	var coord byte
	x, y := s[0]-'0', s[2]-'0'
	switch direction {
	case 1:
		if x != 0 {
			coord = x - 1
			res = strconv.Itoa(int(coord)) + s[1:]
		}
	case 2:
		if x != byte(m-1) {
			coord = x + 1
			res = strconv.Itoa(int(coord)) + s[1:]
		}
	case 3:
		if y != 0 {
			coord = y - 1
			res = s[:2] + strconv.Itoa(int(coord))
		}
	case 4:
		if y != byte(n-1) {
			coord = y + 1
			res = s[:2] + strconv.Itoa(int(coord))
		}
	}
	return res
}
