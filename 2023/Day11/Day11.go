package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	len_row := 0
	grid := make([]string, 0)
	
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		len_row = len(text)
		grid = append(grid, text)
		row++
			
	}
	//fmt.Println(grid)
	fmt.Println(len_row, row)
	
	row_dup_grid := make([]string, 0)

	for r := 0; r < row; r++ {
		flag := false
		for c := 0; c < len_row; c++ {
			if grid[r][c] == '#' {
				flag = true
				break
			}
		}
		if flag {
			row_dup_grid = append(row_dup_grid, grid[r])
		} else {
			row_dup_grid = append(row_dup_grid, grid[r])
			row_dup_grid = append(row_dup_grid, grid[r])	
		}

	}
	//fmt.Println(row_dup_grid)

	//col_row_dup_grid := make([]string, 0)
	num_empty_cols := 0
	col_arr := make([]int, 0)
	for c := 0; c < len_row; c++ {
		flag := false
		for r := 0; r < row; r++ {
			if grid[r][c] == '#' {
				flag = true
				break
			}
		}
		if !flag {
			num_empty_cols++
			col_arr = append(col_arr, c)
		}
	}
	col_row_dup_grid := make([]string, len(row_dup_grid))
	for i := 0; i < len(row_dup_grid); i++ {
		new_row := ""
		col_index := 0
		for j := 0; j < len(row_dup_grid[i]); j++ {
			if col_index < len(col_arr) && j == col_arr[col_index] {
				col_index++
				new_row += string(row_dup_grid[i][j])
			}
			new_row += string(row_dup_grid[i][j])
		}
		col_row_dup_grid[i] = new_row
	}
	fmt.Println(col_row_dup_grid)

	all_pos := make([][]int, 0)
	for r := 0; r < len(col_row_dup_grid); r++ {
		for c := 0; c < len(col_row_dup_grid[r]); c++ {
			if col_row_dup_grid[r][c] == '#' {
				all_pos = append(all_pos, []int{r, c})
			}
		}
	}

	//fmt.Println(all_pos)
	sum := 0
	for i := 0; i < len(all_pos); i++ {
		for j := i + 1; j < len(all_pos); j++ {
			sum += abs(all_pos[i][0] - all_pos[j][0]) + abs(all_pos[i][1] - all_pos[j][1])
		}
	}
	fmt.Println(sum)

}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	len_row := 0
	grid := make([]string, 0)
	
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		len_row = len(text)
		grid = append(grid, text)
		row++
			
	}
	//fmt.Println(grid)
	//fmt.Println(len_row, row)
	row_arr := make([]int, 0)
	
	//row_dup_grid := make([]string, 0)

	for r := 0; r < row; r++ {
		flag := false
		for c := 0; c < len_row; c++ {
			if grid[r][c] == '#' {
				flag = true
				break
			}
		}
		if !flag {
			row_arr = append(row_arr, r)
		}

	}
	//fmt.Println(row_dup_grid)

	col_arr := make([]int, 0)
	for c := 0; c < len_row; c++ {
		flag := false
		for r := 0; r < row; r++ {
			if grid[r][c] == '#' {
				flag = true
				break
			}
		}
		if !flag {
			col_arr = append(col_arr, c)
		}
	}
	

	all_pos := make([][]int, 0)
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '#' {
				all_pos = append(all_pos, []int{r, c})
			}
		}
	}

	//fmt.Println(all_pos)
	sum := 0
	for i := 0; i < len(all_pos); i++ {
		for j := i + 1; j < len(all_pos); j++ {
			var low_r, high_r, low_c, high_c int
			if all_pos[i][0] < all_pos[j][0] {
				low_r = all_pos[i][0]
				high_r = all_pos[j][0]
			} else {
				low_r = all_pos[j][0]
				high_r = all_pos[i][0]
			}
			if all_pos[i][1] < all_pos[j][1] {
				low_c = all_pos[i][1]
				high_c = all_pos[j][1]
			} else {
				low_c = all_pos[j][1]
				high_c = all_pos[i][1]
			}
			val := high_r - low_r + high_c - low_c
			//sum += high_r - low_r + high_c - low_c
			for k := low_r + 1; k < high_r; k++ {
				for l := 0; l < len(row_arr); l++ {
					if row_arr[l] == k {
						val += 999999
					}
				}
			}
			for k := low_c + 1; k < high_c; k++ {
				for l := 0; l < len(col_arr); l++ {
					if col_arr[l] == k {
						val += 999999
					}
				}
			}
			sum += val
		}
	}
	fmt.Println(sum)
}

func main() {
	//part1()
	part2()
}