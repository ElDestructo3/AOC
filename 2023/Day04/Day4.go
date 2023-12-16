package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0


	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		full_nums := strings.Split(text, ": ")
		nums := strings.Split(full_nums[1], " | ")
		win_num_str := strings.Split(nums[0], " ")
		given_num_str := strings.Split(nums[1], " ")

		var win_nums []int
		var given_nums []int

		for i := 0; i < len(win_num_str); i++ {
			if len(win_num_str[i]) > 0 {
				num, _ := strconv.Atoi(win_num_str[i])
				win_nums = append(win_nums, num)
			}
		}
		for i := 0; i < len(given_num_str); i++ {
			if len(given_num_str[i]) > 0 {
				num, _ := strconv.Atoi(given_num_str[i])
				given_nums = append(given_nums, num)
			}
		}


		pow_two := 0
		for i := 0; i < len(win_nums); i++ {
			for j := 0; j < len(given_nums); j++ {
				int_win := win_nums[i]
				int_given := given_nums[j]
				if int_win == int_given {
					if pow_two == 0 {
						pow_two = 1
					} else {
						pow_two *= 2
					}
					break
				}
			}
		}
		sum += pow_two
	}
	fmt.Println(sum)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	var num_matches [194]int
	for i := 0; i < 194; i++ {
		num_matches[i] = 1
	}
	index := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		full_nums := strings.Split(text, ": ")
		nums := strings.Split(full_nums[1], " | ")
		win_num_str := strings.Split(nums[0], " ")
		given_num_str := strings.Split(nums[1], " ")

		var win_nums []int
		var given_nums []int

		for i := 0; i < len(win_num_str); i++ {
			if len(win_num_str[i]) > 0 {
				num, _ := strconv.Atoi(win_num_str[i])
				win_nums = append(win_nums, num)
			}
		}
		for i := 0; i < len(given_num_str); i++ {
			if len(given_num_str[i]) > 0 {
				num, _ := strconv.Atoi(given_num_str[i])
				given_nums = append(given_nums, num)
			}
		}
		count := 0
		for i := 0; i < len(win_nums); i++ {
			for j := 0; j < len(given_nums); j++ {
				int_win := win_nums[i]
				int_given := given_nums[j]
				if int_win == int_given {
					count++
				}
			}
		}
		for i := 1; i <= count; i++ {
			if i + index < 194 {
				num_matches[i + index] += num_matches[index]
			} else {
				break
			}
		}
		sum += num_matches[index]
		index++
		
	}
	fmt.Println(sum)
}

func main() {
	part2()

} 