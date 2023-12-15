package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hash(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
		sum = (sum * 17) % 256
	}
	return sum
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := strings.Split(text, ",")
		for j:= 0; j < len(arr); j++ {
			sum += hash(arr[j])
		}
	}
	fmt.Println(sum)

}

type entry struct {
	id string
	val int
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	hash_map := make([][]entry, 256)
	arr := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr = strings.Split(text, ",")
	}
	for i := 0; i < len(arr); i++ {
		ind := -1
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == '=' || arr[i][j] == '-' {
				ind = j
			}
		}
		hash_val := hash(arr[i][0:ind])
		if arr[i][ind] == '=' {
			val := int(arr[i][len(arr[i]) - 1] - '0')
			new_entry := entry{arr[i][0:ind], val}
			found := false
			for j := 0; j < len(hash_map[hash_val]); j++ {
				if hash_map[hash_val][j].id == arr[i][0:ind] {
					hash_map[hash_val][j].val = val
					found = true
				}
			}
			if !found {
				hash_map[hash_val] = append(hash_map[hash_val], new_entry)
			}	
		} else {
			for j := 0; j < len(hash_map[hash_val]); j++ {
				if hash_map[hash_val][j].id == arr[i][0:ind] {
					new_list := hash_map[hash_val][0:j]
					new_list = append(new_list, hash_map[hash_val][j+1:]...)
					hash_map[hash_val] = new_list
				}
			}
		}
		for i := 0; i < 256; i++ {
			if len(hash_map[i]) > 0 {
				fmt.Println(i, hash_map[i])
			}
		}
		fmt.Println()

	}
	for i := 0; i < 256; i++ {
		for j := 0; j < len(hash_map[i]); j++ {
			sum += hash_map[i][j].val * (i + 1) * (j + 1)
		}
	}
	fmt.Println(sum)



}

func main() {
	//part1()
	part2()
}