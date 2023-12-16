package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		for i := 0; i < len(text); i++ {
			if text[i] >= '0' && text[i] <= '9' {
				sum += 10 * int(text[i] - '0')
				break
			}
		}
		for i := len(text) - 1; i >= 0; i-- {
			if text[i] >= '0' && text[i] <= '9' {
				sum += int(text[i] - '0')
				break
			}
		}
	}
	fmt.Println(sum)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	num_array := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}


	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		for i := 0; i < len(text); i++ {
			if text[i] >= '0' && text[i] <= '9' {
				sum += 10 * int(text[i] - '0')
				break
			}
			
			flag := 0
			for j := 0; j < len(num_array); j++ {
				len_num := len(num_array[j])
				if i + len_num > len(text) {
					continue
				}
				if text[i:i+len_num] == num_array[j] {
					sum += 10 * (j+1)
					flag = 1
					break
				}
			}
			if flag == 1 {
				break
			} 

		}
		for i := len(text) - 1; i >= 0; i-- {
			if text[i] >= '0' && text[i] <= '9' {
				sum += int(text[i] - '0')
				break
			}
			flag := 0
			for j := 0; j < len(num_array); j++ {
				len_num := len(num_array[j])
				if i - len_num + 1 < 0 {
					continue
				}
				if text[i-len_num+1:i+1] == num_array[j] {
					sum += (j+1)
					flag = 1
					break
				}
			}
			if flag == 1 {
				break
			}
		}
		fmt.Println(sum)
	}
	fmt.Println(sum)
}

func main() {
	//part1()
	part2()

} 