package main

import (
	"bufio"
	"fmt"
	"os"
)
// add 1 extra empty line at the end of input file

func shift_north(grid [][]rune) [][]rune {
	new_grid := grid
	for i := 0; i < len(grid[0]); i++ {
		last_hash := -1
		count := 0
		//fmt.Println(i)
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '#' {
				//fmt.Println(j, count)
				for k := 0; k < count; k++ {
					new_grid[last_hash + 1 + k][i] = 'O'
				}
				for k := last_hash + 1 + count; k < j; k++ {
					new_grid[k][i] = '.'
				}
				last_hash = j
				count = 0
			} else if grid[j][i] == 'O' {
				count++
			}
		}
		if count > 0 {
			for k := 0; k < count; k++ {
				new_grid[last_hash + 1 + k][i] = 'O'
			}
			for k := last_hash + 1 + count; k < len(grid); k++ {
				new_grid[k][i] = '.'
			}
		}

	}
	return new_grid
}

func shift_south(grid [][]rune) [][]rune {
	new_grid := grid
	for i := 0; i < len(grid[0]); i++ {
		last_hash := len(grid)
		count := 0
		//fmt.Println(i)
		for j := len(grid) - 1; j >= 0; j-- {
			if grid[j][i] == '#' {
				//fmt.Println(j, count)
				for k := 0; k < count; k++ {
					new_grid[last_hash - 1 - k][i] = 'O'
				}
				for k := last_hash - 1 - count; k > j; k-- {
					new_grid[k][i] = '.'
				}
				last_hash = j
				count = 0
			} else if grid[j][i] == 'O' {
				count++
			}
		}
		if count > 0 {
			for k := 0; k < count; k++ {
				new_grid[last_hash - 1 - k][i] = 'O'
			}
			for k := last_hash - 1 - count; k >= 0; k-- {
				new_grid[k][i] = '.'
			}
		}

	}
	return new_grid
}

func shift_west(grid [][]rune) [][]rune {
	new_grid := grid
	for i := 0; i < len(grid); i++ {
		last_hash := -1
		count := 0
		//fmt.Println(i)
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '#' {
				//fmt.Println(j, count)
				for k := 0; k < count; k++ {
					new_grid[i][last_hash + 1 + k] = 'O'
				}
				for k := last_hash + 1 + count; k < j; k++ {
					new_grid[i][k] = '.'
				}
				last_hash = j
				count = 0
			} else if grid[i][j] == 'O' {
				count++
			}
		}
		if count > 0 {
			for k := 0; k < count; k++ {
				new_grid[i][last_hash + 1 + k] = 'O'
			}
			for k := last_hash + 1 + count; k < len(grid[0]); k++ {
				new_grid[i][k] = '.'
			}
		}

	}
	return new_grid
}

func shift_east(grid [][]rune) [][]rune {
	new_grid := grid
	for i := 0; i < len(grid); i++ {
		last_hash := len(grid[0])
		count := 0
		//fmt.Println(i)
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if grid[i][j] == '#' {
				//fmt.Println(j, count)
				for k := 0; k < count; k++ {
					new_grid[i][last_hash - 1 - k] = 'O'
				}
				for k := last_hash - 1 - count; k > j; k-- {
					new_grid[i][k] = '.'
				}
				last_hash = j
				count = 0
			} else if grid[i][j] == 'O' {
				count++
			}
		}
		if count > 0 {
			for k := 0; k < count; k++ {
				new_grid[i][last_hash - 1 - k] = 'O'
			}
			for k := last_hash - 1 - count; k >= 0; k-- {
				new_grid[i][k] = '.'
			}
		}

	}
	return new_grid
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	grid := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		grid = append(grid, text)
		
	}
	//new_grid := make([]string, len(grid))
	// for i := 0; i < len(grid); i++ {
	// 	new_grid[i] = grid[i]
	// }
	for i := 0; i < len(grid[0]); i++ {
		last_hash := -1
		count := 0
		//fmt.Println(i)
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '#' {
				//fmt.Println(j, count)
				for k := 0; k < count; k++ {
					
					sum += len(grid) - (last_hash + 1 + k)
				}
				last_hash = j
				count = 0
			} else if grid[j][i] == 'O' {
				count++
			}
		}
		if count > 0 {
			for k := 0; k < count; k++ {
				sum += len(grid) - (last_hash + 1 + k)
			}
		}

	}
	fmt.Println(sum)
	

}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]rune, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		grid = append(grid, []rune(text))
		
	}
	period := -1
	cycles := 0
	period_start := -1
	arr := make([][][]rune, 0)
	for cycles < 1000 {
		grid_copy := make([][]rune, len(grid))
		for i := 0; i < len(grid); i++ {
			grid_copy[i] = make([]rune, len(grid[0]))
			for j := 0; j < len(grid[0]); j++ {
				grid_copy[i][j] = grid[i][j]
			}
		}
		grid_copy = shift_north(grid_copy)
		grid_copy = shift_west(grid_copy)
		grid_copy = shift_south(grid_copy)
		grid_copy = shift_east(grid_copy)


		cycles++
		for x := 0; x < len(arr); x++ {
			flag := true
			for i := 0; i < len(grid_copy); i++ {
				for j := 0; j < len(grid_copy[0]); j++ {
					if grid_copy[i][j] != arr[x][i][j] {
						flag = false
						break
					}
					if i == len(grid_copy) - 1 && j == len(grid_copy[0]) - 1 {
						period_start = x
						period = cycles - x - 1
					}
				}
				if !flag {
					break
				}
			}
		}
		if period != -1 {
			break
		}
		arr = append(arr, grid_copy)
		grid = grid_copy

	}
	num_iterations := 1000000000
	num_iterations -= period_start
	num_iterations %= period
	req_grid := arr[period_start + num_iterations - 1]
	sum := 0
	for i := 0; i < len(req_grid); i++ {
		for j := 0; j < len(req_grid[0]); j++ {
			if req_grid[i][j] == 'O' {
				sum += len(req_grid) - i
			}
		}
	}
	fmt.Println(sum)

}

func main() {
	//part1()
	part2()
}