package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

	

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	rules := make(map[string][]string)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		ind := -1
		for i := 0; i < len(text); i++ {
			if text[i] == '{' {
				ind = i
				break
			}
		}
		rule_name := text[:ind]
		rule_list := strings.Split(text[ind + 1:len(text) - 1], ",")
		rules[rule_name] = rule_list
		
	}
	sum := 0
	for scanner.Scan() {
		text:= scanner.Text()
		if text == "" {
			break
		}
		arr := strings.Split(text, ",")
		x := 0
		m := 0
		a := 0
		s := 0
		arr2 := make([]int, 4)
		for i := 0; i < 4; i++ {
			val := 0
			add := 0
			if i == 0 {
				add = 1
			}
			sub := 0
			if i == 3 {
				sub = 1
			}
			for j := 2 + add; j < len(arr[i]) - sub; j++ {
				val = val * 10 + int(arr[i][j]) - int('0')
			}
			arr2[i] = val
		}
		x = arr2[0]
		m = arr2[1]
		a = arr2[2]
		s = arr2[3]

		mp := make(map[byte]int)
		mp['x'] = x
		mp['m'] = m
		mp['a'] = a
		mp['s'] = s

		curr_name := "in"
		is_complete := false
		for !is_complete {
			curr_rule := rules[curr_name]
			for i := 0; i < len(curr_rule); i++ {
				has_colon := false
				ind_col := -1
				for j := 0; j < len(curr_rule[i]); j++ {
					if curr_rule[i][j] == ':' {
						has_colon = true
						ind_col = j
						break
					}
				}
				if !has_colon {
					if curr_rule[i]== "R" {
						is_complete = true
						break
					}
					if curr_rule[i] == "A" {
						sum += x + m + a + s
						is_complete = true
						break
					}
					curr_name = curr_rule[i]
					break
				}
				var_name := curr_rule[i][0]
				is_greater := false
				if curr_rule[i][1] == '>' {
					is_greater = true
				}
				val := 0
				for j := 2; j < ind_col; j++ {
					val = val * 10 + int(curr_rule[i][j]) - int('0')
				}
				if is_greater {
					if mp[var_name] > val {
						curr_name = curr_rule[i][ind_col + 1:]
						break
					}
				} else {
					if mp[var_name] < val {
						curr_name = curr_rule[i][ind_col + 1:]
						break
					}
				}
			}
			if curr_name == "A" {
				is_complete = true
				sum += x + m + a + s
			}
			if curr_name == "R" {
				is_complete = true
			}
		}
	}
	fmt.Println(sum)

}

type bounds struct {
	name string
	bounds [8]int
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	rules := make(map[string][]string)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		ind := -1
		for i := 0; i < len(text); i++ {
			if text[i] == '{' {
				ind = i
				break
			}
		}
		rule_name := text[:ind]
		rule_list := strings.Split(text[ind + 1:len(text) - 1], ",")
		rules[rule_name] = rule_list
		
	}
	sum := 0
	//arr:= []interface{}{"in", 1, 4000, 1, 4000, 1, 4000, 1, 4000} //lower bound, upper bound for 4 variables
	arr_t := make([]bounds, 1)
	arr_t[0].name = "in"
	arr_t[0].bounds = [8]int{1, 4000, 1, 4000, 1, 4000, 1, 4000}
	queue := make([]bounds, 1)
	queue[0] = arr_t[0]
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		//fmt.Println(item)
		curr_name := item.name
		if curr_name == "R" {
			continue
		}
		num := 1
		if curr_name == "A" {
			for i := 0; i < len(item.bounds); i += 2 {
				num *= item.bounds[i + 1] - item.bounds[i] + 1
				if item.bounds[i] > item.bounds[i + 1] {
					num = 0
					break
				}
			}
			//fmt.Println(num)
			sum += num
			continue
		}
		curr_rule := rules[curr_name]
		//fmt.Println(curr_rule)
		for i := 0; i < len(curr_rule); i++ {
			has_colon := false
			ind_col := -1
			for j := 0; j < len(curr_rule[i]); j++ {
				if curr_rule[i][j] == ':' {
					has_colon = true
					ind_col = j
					break
				}
			}
			if !has_colon {
				new_item := item
				new_item.name = curr_rule[i]
				if curr_rule[i] == "R" { 
					continue
				}
				queue = append(queue, new_item)
				continue
			}
			var_name := curr_rule[i][0]
			is_greater := false
			if curr_rule[i][1] == '>' {
				is_greater = true
			}
			val := 0
			for j := 2; j < ind_col; j++ {
				val = val * 10 + int(curr_rule[i][j]) - int('0')
			}
			new_item := item
			new_item.name = curr_rule[i][ind_col + 1:]
			if is_greater {
				if var_name == 'x' {
					new_item.bounds[0] = max(new_item.bounds[0], val + 1)
					item.bounds[1] = min(item.bounds[1], val)
				} else if var_name == 'm' {
					new_item.bounds[2] = max(new_item.bounds[2], val + 1)
					item.bounds[3] = min(item.bounds[3], val)
				} else if var_name == 'a' {
					new_item.bounds[4] = max(new_item.bounds[4], val + 1)
					item.bounds[5] = min(item.bounds[5], val)
				} else if var_name == 's' {
					new_item.bounds[6] = max(new_item.bounds[6], val + 1)
					item.bounds[7] = min(item.bounds[7], val)
				}
			} else {
				if var_name == 'x' {
					new_item.bounds[1] = min(new_item.bounds[1], val - 1)
					item.bounds[0] = max(item.bounds[0], val)
				} else if var_name == 'm' {
					new_item.bounds[3] = min(new_item.bounds[3], val - 1)
					item.bounds[2] = max(item.bounds[2], val)
				} else if var_name == 'a' {
					new_item.bounds[5] = min(new_item.bounds[5], val - 1)
					item.bounds[4] = max(item.bounds[4], val)
				} else if var_name == 's' {
					new_item.bounds[7] = min(new_item.bounds[7], val - 1)
					item.bounds[6] = max(item.bounds[6], val)
				}
			}
			queue = append(queue, new_item)
		}
	}


	
	fmt.Println(sum)
	


}

func main() {
	//part1()
	part2()
}