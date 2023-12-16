package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b % a, a)
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	mp_index := make(map[string]int)
	mp_left := make(map[string]string)
	mp_right := make(map[string]string)
	mp_reverse := make(map[int]string)
	start := -1
	end := -1
	row := 0
	tot := -1
	var moves [400]int
	for scanner.Scan() {
		text := scanner.Text()
		row++
		if row == 1 {
			tot = len(text)
			for ind, c := range text {
				if c == 'R' {
					moves[ind] = 1
				} else if c == 'L' {
					moves[ind] = 0
				}
			}
			continue
		}
		if row == 2 {
			continue
		}
		if text == "" {
			break
		}
		//fmt.Println(text)
		str1 := text[0:3]
		str_left := text[7:10]
		str_right := text[12:15]
		if _, ok := mp_index[str1]; !ok {
			mp_index[str1] = count
			mp_left[str1] = str_left
			mp_right[str1] = str_right
			mp_reverse[count] = str1
			if str1 == "AAA" {
				start = count
			}
			if str1 == "ZZZ" {
				end = count
			}
			count++
		}
			
	}
	//fmt.Println(mp_index)
	steps := 0
	curr := start
	ind := 0
	for {
		if curr == end {
			break
		}
		move := moves[ind]
		if move == 0 {
			curr = mp_index[mp_left[mp_reverse[curr]]]
		} else {
			curr = mp_index[mp_right[mp_reverse[curr]]]
		}
		ind = (ind + 1) % tot
		steps++
	}
	fmt.Println(steps)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	mp_index := make(map[string]int)
	mp_left := make(map[string]string)
	mp_right := make(map[string]string)
	mp_reverse := make(map[int]string)
	var start []int
	var end []int
	row := 0
	tot := -1
	var moves [400]int
	for scanner.Scan() {
		text := scanner.Text()
		row++
		if row == 1 {
			tot = len(text)
			for ind, c := range text {
				if c == 'R' {
					moves[ind] = 1
				} else if c == 'L' {
					moves[ind] = 0
				}
			}
			continue
		}
		if row == 2 {
			continue
		}
		if text == "" {
			break
		}
		//fmt.Println(text)
		str1 := text[0:3]
		str_left := text[7:10]
		str_right := text[12:15]
		if _, ok := mp_index[str1]; !ok {
			mp_index[str1] = count
			mp_left[str1] = str_left
			mp_right[str1] = str_right
			mp_reverse[count] = str1
			if str1[2] == 'A' {
				start = append(start, count)
			}
			if str1[2] == 'Z' {
				end = append(end, count)
			}
			count++
		}
			
	}
	//fmt.Println(mp_index)
	var steps []int
	for _, curr := range start {
		ind := 0
		curr_steps := 0
		flag := false
		for {
			for _, e := range end {
				if curr == e {
					steps = append(steps, curr_steps)
					flag = true
					break
				}
			}
			if flag {
				break
			}

			move := moves[ind]
			if move == 0 {
				curr = mp_index[mp_left[mp_reverse[curr]]]
			} else {
				curr = mp_index[mp_right[mp_reverse[curr]]]
			}
			ind = (ind + 1) % tot
			curr_steps++
		}
	}
	lcm := steps[0]
	for i := 1; i < len(steps); i++ {
		lcm = lcm * steps[i] / gcd(lcm, steps[i])
	}
	fmt.Println(lcm)
}

func main() {
	//part1()
	part2()
}