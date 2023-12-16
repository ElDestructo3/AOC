package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func card_rank(card string) int {
	var mp map[byte]int
	mp = make(map[byte]int)
	for i := 0; i < len(card); i++ {
		mp[card[i]]++
	}
	if len(mp) == 1 {
		return 7
	}
	if len(mp) == 2 {
		for _, v := range mp {
			if v == 4 || v == 1 {
				return 6
			} else {
				return 5
			}
		}
	}
	var arr []int
	for _, v := range mp {
		arr = append(arr, v)
	}
	sort.Ints(arr)
	if len(arr) == 3 {
		if arr[2] == 3 {
			return 4
		} else {
			return 3
		}
	}
	if len(arr) == 4 {
		return 2
	}
	return 1


}

func card_rank_part2(card string) int {
	index := 0
	for _, character := range card {
		if character == 'J' {
			index++
		}
	}
	if index == 0 {
		return card_rank(card)
	} else {
		var mp map[byte]int
		mp = make(map[byte]int)
		for _, character := range card {
			mp[byte(character)]++
		}
		max_freq := 0
		max_char := byte(0)
		for k, v := range mp {
			if v > max_freq && k != byte('J'){
				max_freq = v
				max_char = k
			}
		}
		var new_card string
		for _, character := range card {
			if character == rune('J') {
				new_card += string(max_char)
			} else {
				new_card += string(character)
			}
		}
		return card_rank(new_card)
		
	}
}

func is_better(card1 string, card2 string) bool {

	
	if card_rank(card1) > card_rank(card2) {
		return true
	} else if card_rank(card1) < card_rank(card2) {
		return false
	} else {
		for i := 0; i < len(card1); i++ {
			var val1 int
			var val2 int
			if card1[i] >= '0' && card1[i] <= '9' {
				val1 = int(card1[i] - '0')
			} else {
				switch card1[i] {
					case 'T':
						val1 = 10
					case 'J':
						val1 = 11
					case 'Q':
						val1 = 12
					case 'K':
						val1 = 13
					case 'A':
						val1 = 14
					}
			}
			if card2[i] >= '0' && card2[i] <= '9' {
				val2 = int(card2[i] - '0')
			} else {
				switch card2[i] {
					case 'T':
						val2 = 10
					case 'J':
						val2 = 11
					case 'Q':
						val2 = 12
					case 'K':
						val2 = 13
					case 'A':
						val2 = 14
				}
			}
			if val1 > val2 {
				return true
			} else if val1 < val2 {
				return false
			}
		}
	}
	return false
}

func is_better_part2(card1 string, card2 string) bool {

	if card_rank_part2(card1) > card_rank_part2(card2) {
		return true
	} else if card_rank_part2(card1) < card_rank_part2(card2) {
		return false
	} else {
		var arr1 []int
		var arr2 []int
		for _, character := range card1 {
			switch character {
				case 'T':
					arr1 = append(arr1, 10)
				case 'J':
					arr1 = append(arr1, 0)
				case 'Q':
					arr1 = append(arr1, 11)
				case 'K':
					arr1 = append(arr1, 12)
				case 'A':
					arr1 = append(arr1, 13)
				default:
					arr1 = append(arr1, int(character - '0'))
				}
		}
		for _, character := range card2 {
			switch character {
				case 'T':
					arr2 = append(arr2, 10)
				case 'J':
					arr2 = append(arr2, 0)
				case 'Q':
					arr2 = append(arr2, 11)
				case 'K':
					arr2 = append(arr2, 12)
				case 'A':
					arr2 = append(arr2, 13)
				default:
					arr2 = append(arr2, int(character - '0'))
				}
		}
		print(arr1, arr2)
		for i := 0; i < len(arr1); i++ {
			if arr1[i] > arr2[i] {
				return true
			} else if arr1[i] < arr2[i] {
				return false
			}
		}
		

	}
	return false
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	var cards []string
	var bids []int
	

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := strings.Split(text, " ")
		cards = append(cards, arr[0])
		int_val, _ := strconv.Atoi(arr[1])
		bids = append(bids, int_val)
	}
	
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if is_better(cards[i], cards[j]) {
				cards[i], cards[j] = cards[j], cards[i]
				bids[i], bids[j] = bids[j], bids[i]
			}
		}
	}
	for i := 0; i < len(cards); i++ {
		fmt.Println(cards[i], card_rank(cards[i]))
		sum += (i + 1) * bids[i]
	}
	fmt.Println(sum)

}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	var cards []string
	var bids []int
	

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := strings.Split(text, " ")
		cards = append(cards, arr[0])
		int_val, _ := strconv.Atoi(arr[1])
		bids = append(bids, int_val)
	}
	
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if is_better_part2(cards[i], cards[j]) {
				cards[i], cards[j] = cards[j], cards[i]
				bids[i], bids[j] = bids[j], bids[i]
			}
		}
	}
	for i := 0; i < len(cards); i++ {
		fmt.Println(cards[i], card_rank(cards[i]))
		sum += (i + 1) * bids[i]
	}
	fmt.Println(sum)
}

func main() {
	part2()

} 