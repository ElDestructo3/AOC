package main

import (
	"bufio"
	"fmt"
	"os"
)

func twotoone(x, y int) int {
	return (x + 1) * 10000 + (y + 1)
}

func onetotwo(n int) (int, int) {
	return n / 10000 - 1, n % 10000 - 1
}

func energise(grid []string, start int, par int) int {
	count := 0

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{start, par})
	max_one := twotoone(len(grid) - 1, len(grid[0]) - 1)
	visited := make([][2]int, 0)
	parents := make([]int, max_one + 1)
	for i := 0; i < len(parents); i++ {
		parents[i] = -1
	}
	parents[start] = twotoone(0, -1)
	iters := 0
	for len(queue) > 0 && iters < 10000000000 {
		iters++
		pos := queue[0][0]
		par := queue[0][1]
		queue = queue[1:]
		x, y := onetotwo(pos)
		par_x, par_y := onetotwo(par)
		new_x := 2 * x - par_x
		new_y := 2 * y - par_y
		dir_x := x - par_x 
		dir_y := y - par_y
		next_pos := make([][2]int, 0)
		if grid[x][y] == '.' {

			if new_x < 0 || new_x >= len(grid) || new_y < 0 || new_y >= len(grid[0]) {
				continue
			}
			next_pos = append(next_pos, [2]int{new_x, new_y})
		}
		if grid[x][y] == '|' {
			if dir_x != 0 {
				if new_x < 0 || new_x >= len(grid) || new_y < 0 || new_y >= len(grid[0]) {
					continue
				}

				next_pos = append(next_pos, [2]int{new_x, new_y})
			} else {
				if x + 1 < len(grid){ 
					next_pos = append(next_pos, [2]int{x + 1, y})
				}
				if x - 1 >= 0 {
					next_pos = append(next_pos, [2]int{x - 1, y})
				}
			}
		}
		if grid[x][y] == '-' {
			if dir_y != 0 {
				if new_x < 0 || new_x >= len(grid) || new_y < 0 || new_y >= len(grid[0]) {
					continue
				}
				next_pos = append(next_pos, [2]int{new_x, new_y})

			} else {
				if y + 1 < len(grid[0]) {
					next_pos = append(next_pos, [2]int{x, y + 1})
				}
				if y - 1 >= 0 {
					next_pos = append(next_pos, [2]int{x, y - 1})
				}
			}
		}
		if grid[x][y] == '/' {
			if dir_x < 0 {
				if y + 1 < len(grid[0]) {
					next_pos = append(next_pos, [2]int{x, y + 1})
				}
			} else if dir_x > 0 {
				if y - 1 >= 0 {
					next_pos = append(next_pos, [2]int{x, y - 1})
				}
			} else {
				if dir_y > 0 {
					if x - 1 >= 0 {
						next_pos = append(next_pos, [2]int{x - 1, y})
					}
				} else {
					if x + 1 < len(grid) {
						next_pos = append(next_pos, [2]int{x + 1, y})
					}
				}
			}
		}
		if grid[x][y] == '\\' {
			if dir_x < 0 {
				if y - 1 >= 0 {
					next_pos = append(next_pos, [2]int{x, y - 1})
				}
			} else if dir_x > 0 {
				if y + 1 < len(grid[0]) {
					next_pos = append(next_pos, [2]int{x, y + 1})
				}
			} else {
				if dir_y > 0 {
					if x + 1 < len(grid) {
						next_pos = append(next_pos, [2]int{x + 1, y})
					}
				} else {
					if x - 1 >= 0 {
						next_pos = append(next_pos, [2]int{x - 1, y})
					}
				}
			}
		}
		for i := 0; i < len(next_pos); i++ {
			flag := true
			for j := 0; j < len(visited); j++ {
				if visited[j][1] == pos && visited[j][0] == twotoone(next_pos[i][0], next_pos[i][1]) {
					flag = false
					break
				}
			}
			if flag {
				queue = append(queue, [2]int{twotoone(next_pos[i][0], next_pos[i][1]), pos})
				parents[twotoone(next_pos[i][0], next_pos[i][1])] = pos
				visited = append(visited, [2]int{twotoone(next_pos[i][0], next_pos[i][1]), pos})
			}
		}
	}
	for i := 0; i < len(parents); i++ {
		if parents[i] != -1 {
			count++
		}
	}
	return count
}
	

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		grid = append(grid, text)
	}
	fmt.Println(energise(grid, twotoone(0, 0), twotoone(0, -1)))


}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		grid = append(grid, text)
	}
	max_val := -1
	for i := 0; i < len(grid); i++ {
		max_val = max(max_val, energise(grid, twotoone(i, 0), twotoone(i, -1)))
		max_val = max(max_val, energise(grid, twotoone(i, len(grid[0]) - 1), twotoone(i, len(grid[0]))))
	}
	for i := 0; i < len(grid[0]); i++ {
		max_val = max(max_val, energise(grid, twotoone(0, i), twotoone(-1, i)))
		max_val = max(max_val, energise(grid, twotoone(len(grid) - 1, i), twotoone(len(grid), i)))
	}
	fmt.Println(max_val)




}

func main() {
	//part1()
	part2()
}