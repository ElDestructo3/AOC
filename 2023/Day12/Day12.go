package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)



func solve(s_row string, curr_index int, streak int, streak_arr []int, streak_index int, dp [][][]int) int {
	//fmt.Println(curr_index, streak, streak_index)
	if curr_index == len(s_row) {
		if streak_index == len(streak_arr) || (streak_index == len(streak_arr) - 1 && streak == streak_arr[streak_index]){
			return 1
		} else {
			return 0
		}
	}
	if dp[curr_index][streak][streak_index] != -1 {
		return dp[curr_index][streak][streak_index]
	}
	s1 := 0
	s2 := 0
	if s_row[curr_index] == '#' || s_row[curr_index] == '?' {
		new_streak := streak + 1
		if streak_index < len(streak_arr) && new_streak <= streak_arr[streak_index] {
			s1 = solve(s_row, curr_index + 1, new_streak, streak_arr, streak_index, dp)
		}
	} 
	if s_row[curr_index] == '.' || s_row[curr_index] == '?' {
		if streak_index < len(streak_arr) && streak == streak_arr[streak_index] {
			s2 = solve(s_row, curr_index + 1, 0, streak_arr, streak_index + 1, dp)
		} else if streak == 0 {
			s2 = solve(s_row, curr_index + 1, 0, streak_arr, streak_index, dp)
		}
	}
	dp[curr_index][streak][streak_index] = s1 + s2
	return s1 + s2

}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := strings.Split(text, " ")
		s_row := arr[0]
		s_streak_arr := strings.Split(arr[1], ",")
		streak_arr := make([]int, 0)
		for i := 0; i < len(s_streak_arr); i++ {
			val, _ := strconv.Atoi(s_streak_arr[i])
			streak_arr = append(streak_arr, val)
		}
		//fmt.Println(s_row)
		//fmt.Println(len(s_row))
		//fmt.Println(streak_arr)
		dp := make([][][]int, len(s_row))
		for i := 0; i < len(s_row); i++ {
			dp[i] = make([][]int, len(s_row) + 1)
			for j := 0; j < len(s_row) + 1; j++ {
				dp[i][j] = make([]int, len(streak_arr) + 1)
				for k := 0; k < len(streak_arr) + 1; k++ {
					dp[i][j][k] = -1
				}
			}
		}
		sum += solve(s_row, 0, 0, streak_arr, 0, dp)
		
	}
	fmt.Println(sum)
	

}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := strings.Split(text, " ")
		s_row := arr[0]
		s_streak_arr := strings.Split(arr[1], ",")
		streak_arr := make([]int, 0)
		for i := 0; i < len(s_streak_arr); i++ {
			val, _ := strconv.Atoi(s_streak_arr[i])
			streak_arr = append(streak_arr, val)
		}
		s_row_five := s_row + "?" + s_row + "?" + s_row + "?" + s_row + "?" + s_row
		streak_arr_five := make([]int, 0)
		for i := 0; i < 5; i++ {
			for j := 0; j < len(streak_arr); j++ {
				streak_arr_five = append(streak_arr_five, streak_arr[j])
			}
		}
		dp := make([][][]int, len(s_row_five))
		for i := 0; i < len(s_row_five); i++ {
			dp[i] = make([][]int, len(s_row_five) + 1)
			for j := 0; j < len(s_row_five) + 1; j++ {
				dp[i][j] = make([]int, len(streak_arr_five) + 1)
				for k := 0; k < len(streak_arr_five) + 1; k++ {
					dp[i][j][k] = -1
				}
			}
		}
		a := solve(s_row_five, 0, 0, streak_arr_five, 0, dp)
		sum += a
		
	}
	fmt.Println(sum)
}

func main() {
	//part1()
	part2()
}