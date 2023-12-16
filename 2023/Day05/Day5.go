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
	text_read_stage := 0
	seed_arr := []int{}
	seed_soil := [][]int{}
	soil_fert := [][]int{}
	fert_water := [][]int{}
	water_light := [][]int{}
	light_temp := [][]int{}
	temp_humid := [][]int{}
	humid_loc := [][]int{}

	var mp map[int][][]int
	mp = make(map[int][][]int)

	mp[1] = seed_soil
	mp[2] = soil_fert
	mp[3] = fert_water
	mp[4] = water_light
	mp[5] = light_temp
	mp[6] = temp_humid
	mp[7] = humid_loc

	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Println(text)
		if text == "" {
			text_read_stage++
			continue
		}
		if text_read_stage == 0 {
			seeds := strings.Split(text, ": ")[1]
			seed_split := strings.Split(seeds, " ")
			for _, seed := range seed_split {
				int_val, _ := strconv.Atoi(seed)
				seed_arr = append(seed_arr, int_val)
			}
		} else {
			for scanner.Scan() {
				curr_text := scanner.Text()
				//fmt.Println(curr_text)
				if curr_text == "" {
					text_read_stage++
					break
				}
				arr := strings.Split(curr_text, " ")
				conv_arr := []int{}
				for _, val := range arr {
					int_val, _ := strconv.Atoi(val)
					conv_arr = append(conv_arr, int_val)
				}
				// seed_soil = append(seed_soil, conv_arr)
				mp[text_read_stage] = append(mp[text_read_stage], conv_arr)
			}
		}
		
	}
	ans := 1000000000
	for _, val := range seed_arr {
		loc := val
		for index := 1; index <= 7; index++ {
			//fmt.Println(loc)
			for _, arr := range mp[index] {
				//fmt.Println(arr[1] + arr[2])
				if arr[1] + arr[2] > loc && arr[1] <= loc {
					loc = arr[0] + (loc - arr[1])
					break
				}
			}
		}
		if loc < ans {
			ans = loc
		}
		//fmt.Println(loc)
	}
	fmt.Println(ans)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	text_read_stage := 0
	seed_arr := []int{}
	seed_soil := [][]int{}
	soil_fert := [][]int{}
	fert_water := [][]int{}
	water_light := [][]int{}
	light_temp := [][]int{}
	temp_humid := [][]int{}
	humid_loc := [][]int{}

	var mp map[int][][]int
	mp = make(map[int][][]int)

	mp[1] = seed_soil
	mp[2] = soil_fert
	mp[3] = fert_water
	mp[4] = water_light
	mp[5] = light_temp
	mp[6] = temp_humid
	mp[7] = humid_loc

	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Println(text)
		if text == "" {
			text_read_stage++
			continue
		}
		if text_read_stage == 0 {
			seeds := strings.Split(text, ": ")[1]
			seed_split := strings.Split(seeds, " ")
			for _, seed := range seed_split {
				int_val, _ := strconv.Atoi(seed)
				seed_arr = append(seed_arr, int_val)
			}
		} else {
			for scanner.Scan() {
				curr_text := scanner.Text()
				//fmt.Println(curr_text)
				if curr_text == "" {
					text_read_stage++
					break
				}
				arr := strings.Split(curr_text, " ")
				conv_arr := []int{}
				for _, val := range arr {
					int_val, _ := strconv.Atoi(val)
					conv_arr = append(conv_arr, int_val)
				}
				// seed_soil = append(seed_soil, conv_arr)
				mp[text_read_stage] = append(mp[text_read_stage], conv_arr)
			}
		}
		
	}
	ans := 1000000000
	for i := 0; i < len(seed_arr); i+=2 {
		start := seed_arr[i]
		end := seed_arr[i+1] + start - 1
		for val := start; val <= end; val++ {
			loc := val
			for index := 1; index <= 7; index++ {
				//fmt.Println(loc)
				for _, arr := range mp[index] {
					//fmt.Println(arr[1] + arr[2])
					if arr[1] + arr[2] > loc && arr[1] <= loc {
						loc = arr[0] + (loc - arr[1])
						break
					}
				}
			}
			if loc < ans {
				ans = loc
			}
			
		}
		fmt.Println("done iteration")
		
	}
	fmt.Println(ans)
}

func main() {
	part2()

} 