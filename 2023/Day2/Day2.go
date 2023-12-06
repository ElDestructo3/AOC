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
	max_red := 12
	max_green := 13
	max_blue := 14

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		flag := true
		arr := strings.Split(text, ": ")
		id := strings.Split(arr[0], " ")[1]
		id_num, _ := strconv.Atoi(id)
		sub_games := strings.Split(arr[1], "; ")
		for sub_game := range sub_games {
			sub_game_arr := strings.Split(sub_games[sub_game], ", ")
			for color := 0 ; color < len(sub_game_arr); color++ {
				event := strings.Split(sub_game_arr[color], " ");
				if event[1] == "red"  {
					num_red, _ := strconv.Atoi(event[0])
					if num_red > max_red {
						flag = false
					}
				} else if event[1] == "green" {
					num_green, _ := strconv.Atoi(event[0])
					if num_green > max_green {
						flag = false
					}
				} else if event[1] == "blue" {
					num_blue, _ := strconv.Atoi(event[0])
					if num_blue > max_blue {
						flag = false
					}
				}

			}
				
		}
		if flag {
			sum += id_num
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
		max_red := 0
		max_green := 0
		max_blue := 0
		arr := strings.Split(text, ": ")
		sub_games := strings.Split(arr[1], "; ")
		for sub_game := range sub_games {
			sub_game_arr := strings.Split(sub_games[sub_game], ", ")
			for color := 0 ; color < len(sub_game_arr); color++ {
				event := strings.Split(sub_game_arr[color], " ");
				//fmt.Println(event)
				if event[1] == "red"  {
					num_red, _ := strconv.Atoi(event[0])
					if num_red > max_red {
						max_red = num_red
					}
				} else if event[1] == "green" {
					num_green, _ := strconv.Atoi(event[0])
					if num_green > max_green {
						max_green = num_green
					}
				} else if event[1] == "blue" {
					num_blue, _ := strconv.Atoi(event[0])
					if num_blue > max_blue {
						max_blue = num_blue
					}
				}

			}
				
		}
		sum += max_red * max_green * max_blue

	}
	fmt.Println(sum)
}

func main() {
	part1()
	//part2()
}