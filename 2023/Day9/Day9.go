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

		arr := strings.Split(text, " ")
		int_arr := make([]int, len(arr))
		for i, s := range arr {
			int_arr[i], _ = strconv.Atoi(s)
		}
		var last_val []int
		var last_diff []int
		arr_copy := make([]int, len(int_arr))
		copy(arr_copy, int_arr)
		for {
			flag := true
			new_arr := make([]int, len(arr_copy) - 1)
			for i := 0; i < len(arr_copy) - 1; i++ {
				new_arr[i] = arr_copy[i + 1] - arr_copy[i]
				if new_arr[i] != 0 {
					flag = false
				}
			}
			if flag {
				break
			} else {
				last_val = append(last_val, arr_copy[len(arr_copy) - 1])
				last_diff = append(last_diff, new_arr[len(new_arr) - 1])
				arr_copy = new_arr
			}
		}
		
		if len(last_val) == 0 {
			sum += int_arr[0]
		} else {
			curr := last_val[len(last_val) - 1] + last_diff[len(last_diff) - 1]
			for i := len(last_val) - 2; i >= 0; i-- {
				curr += last_val[i]
			}
			sum += curr
		}

		
			
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
		int_arr := make([]int, len(arr))
		for i, s := range arr {
			int_arr[i], _ = strconv.Atoi(s)
		}
		var last_val []int
		var last_diff []int
		arr_copy := make([]int, len(int_arr))
		copy(arr_copy, int_arr)
		for {
			flag := true
			new_arr := make([]int, len(arr_copy) - 1)
			for i := 0; i < len(arr_copy) - 1; i++ {
				new_arr[i] = arr_copy[i + 1] - arr_copy[i]
				if new_arr[i] != 0 {
					flag = false
				}
			}
			if flag {
				break
			} else {
				last_val = append(last_val, arr_copy[0])
				last_diff = append(last_diff, new_arr[0])
				arr_copy = new_arr
			}
		}
		
		if len(last_val) == 0 {
			sum += int_arr[0]
		} else {
			curr := last_val[len(last_val) - 1] - last_diff[len(last_diff) - 1]
			for i := len(last_val) - 2; i >= 0; i-- {
				curr = last_val[i] - curr
			}
			sum += curr
		}

		
			
	}
	fmt.Println(sum)
}

func main() {
	//part1()
	part2()
}