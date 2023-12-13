package main

import (
	"bufio"
	"fmt"
	"os"
)
// add 1 extra empty line at the end of input file

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	pattern := make([]string, 0)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			new_pattern := make([]string, 0)
			for i := 0; i < len(pattern); i++ {
				if len(pattern[i]) == 0 {
					continue
				}
				new_pattern = append(new_pattern, pattern[i])
			}
			pattern = new_pattern

			num_rows := len(pattern)
			num_cols := len(pattern[0])

			row_sum := 0
			col_sum := 0

			for row_sym := 0; row_sym < num_rows - 1; row_sym++ {
				up := row_sym
				down := row_sym + 1
				flag := true
				for up >= 0 && down < num_rows {
					if pattern[up] != pattern[down] {
						flag = false
						break
					}
					up--
					down++
				}
				if flag {
					row_sum += row_sym + 1
				}
			}

			for col_sym := 0; col_sym < num_cols - 1; col_sym++ {
				left := col_sym
				right := col_sym + 1
				flag := true
				for left >= 0 && right < num_cols {
					for row := 0; row < num_rows; row++ {
						if pattern[row][left] != pattern[row][right] {
							flag = false
							break
						}
					}
					if !flag {
						break
					}
					left--
					right++
				}
				if flag {
					col_sum += col_sym + 1
				}
			}

			sum += 100 * row_sum + col_sum
			pattern = make([]string, 0)
		}
		pattern = append(pattern, text)
		
	}
	fmt.Println(sum)
	

}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	pattern := make([]string, 0)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			new_pattern := make([]string, 0)
			for i := 0; i < len(pattern); i++ {
				if len(pattern[i]) == 0 {
					continue
				}
				new_pattern = append(new_pattern, pattern[i])
			}
			pattern = new_pattern

			num_rows := len(pattern)
			num_cols := len(pattern[0])

			row_sum := 0
			col_sum := 0

			for row_sym := 0; row_sym < num_rows - 1; row_sym++ {
				up := row_sym
				down := row_sym + 1
				chars_diff := 0
				char_pos := make([]int, 2)
				flag := true
				for up >= 0 && down < num_rows {
					if pattern[up] != pattern[down] {
						for i := 0; i < num_cols; i++ {
							if pattern[up][i] != pattern[down][i] {
								chars_diff++
								char_pos[0] = up
								char_pos[1] = i
							}
						}
						if chars_diff > 1 {
							flag = false
							break
						}
					}
					up--
					down++
				}
				if flag && chars_diff == 1 {
					row_sum += row_sym + 1

				}
			}

			for col_sym := 0; col_sym < num_cols - 1; col_sym++ {
				left := col_sym
				right := col_sym + 1
				chars_diff := 0
				char_pos := make([]int, 2)
				flag := true
				for left >= 0 && right < num_cols {
					for row := 0; row < num_rows; row++ {
						if pattern[row][left] != pattern[row][right] {
							chars_diff++
							char_pos[0] = row
							char_pos[1] = left
							if chars_diff > 1 {
								flag = false
								break
							}
						}
					}
					if !flag {
						break
					}
					left--
					right++
				}
				if flag && chars_diff == 1 {
					col_sum += col_sym + 1
				}
			}

			sum += 100 * row_sum + col_sum
			pattern = make([]string, 0)
		}
		pattern = append(pattern, text)
		
	}
	fmt.Println(sum)
	
}

func main() {
	//part1()
	part2()
}