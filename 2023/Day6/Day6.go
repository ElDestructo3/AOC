package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	time_arr := []int{}
	dist_arr := []int{}
	count := 0

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		if text == "" {
			break
		}

		if count == 0 {
			curr_val := 0
			for _, val := range text {
				if val == ' ' {
					fmt.Println(curr_val)
					if curr_val != 0 {
						time_arr = append(time_arr, curr_val)
						curr_val = 0
					}
					continue
				}
				if val >= '0' && val <= '9' {
					curr_val *= 10
					curr_val += int(val - '0')
				}
			}
			if curr_val != 0 {
				time_arr = append(time_arr, curr_val)
			}
			count++

		} else if count == 1 {
			curr_val := 0
			for _, val := range text {
				if val == ' ' {
					if curr_val != 0 {
						dist_arr = append(dist_arr, curr_val)
						curr_val = 0
					}
					continue
				}
				if val >= '0' && val <= '9' {
					curr_val *= 10
					curr_val += int(val - '0')
				}
			}
			if curr_val != 0 {
				dist_arr = append(dist_arr, curr_val)
			}
			count++
		}
	}
	ans := 1
	for i := 0; i < len(time_arr); i++ {
		time := time_arr[i]
		dist := dist_arr[i]
		pos := 0
		for j := 0; j <= time; j++ {
			poss_dist := (time - j) * j
			if poss_dist > dist {
				pos++
			}
		}
		ans *= pos
	}
	fmt.Println(ans)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	time_arr := 0
	dist_arr := 0
	count := 0

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		if text == "" {
			break
		}

		if count == 0 {
			for _, val := range text {
				if val >= '0' && val <= '9' {
					time_arr *= 10
					time_arr += int(val - '0')
				}
			}
			count++

		} else if count == 1 {
			for _, val := range text {
				if val >= '0' && val <= '9' {
					dist_arr *= 10
					dist_arr += int(val - '0')
				}
			}
			count++
		}
	}
	ans := 1
	disc := time_arr * time_arr - 4 * dist_arr
	sqrt_dist := math.Sqrt(float64(disc))
	min_x := math.Ceil((float64(time_arr) - sqrt_dist) / 2)
	max_x := math.Floor((float64(time_arr) + sqrt_dist) / 2)
	fmt.Println(min_x, max_x)
	ans = int(max_x - min_x + 1)
	fmt.Println(ans)
	
}

func main() {
	part2()

} 