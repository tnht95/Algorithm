package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type cell struct {
	y int
	x int
}

var (
	matrix  []string
	lakes   [][]cell
	n, m, k int
	dx      = []int{-1, 0, 1, 0}
	dy      = []int{0, -1, 0, 1}
	visited = [][]bool{}
)

func dfs(i int, j int) {
	isOcean := false
	temp := []cell{}
	stack := []cell{cell{i, j}}
	visited[i][j] = true

	for len(stack) > 0 {
		cur := pop(&stack)
		if onBorder(cur.y, cur.x) {
			isOcean = true
		}
		for d := 0; d < 4; d++ {
			newX := cur.x + dx[d]
			newY := cur.y + dy[d]
			if isValid(newX, newY) && matrix[newY][newX] == '.' && !visited[newY][newX] {
				visited[newY][newX] = true
				stack = append(stack, cell{newY, newX})
				if !isOcean {
					temp = append(temp, cell{newY, newX})
				}
			}
		}
	}
	if !isOcean {
		temp = append(temp, cell{i, j})
		lakes = append(lakes, temp)
	}
}

func pop(s *[]cell) cell {
	n := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return n
}

func onBorder(y int, x int) bool {
	return y == 0 || x == 0 || y == n-1 || x == m-1
}

func isValid(x int, y int) bool {
	return 0 <= x && x < m && 0 <= y && y < n
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d %d %d\n", &n, &m, &k)
	visited = make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	matrix = make([]string, n)
	for i := range matrix {
		fmt.Fscanf(r, "%s\n", &matrix[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && matrix[i][j] == '.' {
				dfs(i, j)
			}
		}
	}

	//sort to get smallest lakes to fill in
	sort.Slice(lakes, func(i, j int) bool { return len(lakes[i]) < len(lakes[j]) })
	cells := 0
	num := len(lakes) - k
	for i := 0; i < num; i++ {
		cells += len(lakes[i])
		for j := range lakes[i] {
			c := lakes[i][j]
			arr := []byte(matrix[c.y])
			arr[c.x] = '*'
			matrix[c.y] = string(arr)
		}
	}

	fmt.Println(cells)
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}
