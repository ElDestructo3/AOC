package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	var grid [1000][1000]int
	for r := 0; r < 1000; r++ {
		for c := 0; c < 1000; c++ {
			grid[r][c] = -1
		}
	}
	row := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		for c := 0; c < len(text); c++ {
			if text[c] >= '0' && text[c] <= '9' {
				num, _ := strconv.Atoi(string(text[c]))
				grid[row][c] = num
			} else if text[c] == '.' {
				grid[row][c] = -1
			} else {
				grid[row][c] = -2
			}
		}
		row++
			
	}
	//fmt.Println(grid)
	for r := 0; r < row; r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] >= 0 {
				//fmt.Println(r, c)
				flag := false
				k := c
				num := 0
				for k < len(grid[r]) && grid[r][k] >= 0 {
					num = num * 10 + grid[r][k]
					k++
				}
				if c > 0 {
					if grid[r][c - 1] == -2 {
						flag = true
					}
				}
				if k < len(grid[r]) {
					if grid[r][k] == -2 {
						flag = true
					}
				}
				if r > 0 {
					x := c - 1
					if x < 0 {
						x = 0	
					} 
					for x <= k && x < len(grid[r - 1]) {
						if grid[r - 1][x] == -2 {
							flag = true					
						}
						x++
					}
				}	

				if r < row - 1 {
					x := c - 1
					if x < 0 {
						x = 0
					} 
					for (x <= k && x < len(grid[r + 1])) {
						if grid[r + 1][x] == -2 {
							flag = true
						}
						x++
					}

				}
				c = k
				if flag {
					sum += num
					//fmt.Println(num)
				}
			}
		}
	}
	fmt.Println(sum)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	var grid [1000][1000]int
	var counts [1000][1000]int
	var prods [1000][1000]int
	for r := 0; r < 1000; r++ {
		for c := 0; c < 1000; c++ {
			grid[r][c] = -1
			counts[r][c] = 0
			prods[r][c] = 1
		}
	}
	row := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		for c := 0; c < len(text); c++ {
			if text[c] >= '0' && text[c] <= '9' {
				num, _ := strconv.Atoi(string(text[c]))
				grid[row][c] = num
			} else if text[c] == '.' {
				grid[row][c] = -1
			} else if text[c] == '*'{ 
				grid[row][c] = -2
			} else {
				grid[row][c] = -3
			}
		}
		row++
			
	}
	//fmt.Println(grid)
	for r := 0; r < row; r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] >= 0 {
				//fmt.Println(r, c)
				k := c
				num := 0
				for k < len(grid[r]) && grid[r][k] >= 0 {
					num = num * 10 + grid[r][k]
					k++
				}
				if c > 0 {
					if grid[r][c - 1] <= -2 {
						if grid[r][c - 1] == -2 {
							//fmt.Println(r, c)
							counts[r][c - 1]++;
							prods[r][c - 1] *= num
						}
					}
				}
				if k < len(grid[r]) {
					if grid[r][k] <= -2{
						if grid[r][k] == -2 {
							//fmt.Println(r, c)
							counts[r][k]++;
							prods[r][k] *= num
						}
					}
				}
				if r > 0 {
					x := c - 1
					if x < 0 {
						x = 0	
					} 
					for x <= k && x < len(grid[r - 1]) {
						if grid[r - 1][x] <= -2 {
							if grid[r - 1][x] == -2 {
								counts[r - 1][x]++;
								prods[r - 1][x] *= num
							}			
						}
						x++
					}
				}	

				if r < row - 1 {
					x := c - 1
					if x < 0 {
						x = 0
					} 
					for (x <= k && x < len(grid[r + 1])) {
						if grid[r + 1][x] <= -2 {
							if grid[r + 1][x] == -2 {
								counts[r + 1][x]++;
								prods[r + 1][x] *= num
							}
						}
						x++
					}

				}
				c = k
			}
		}
	}
	for r := 0; r < row; r++ {
		for c := 0; c < len(grid[r]); c++ {
			if counts[r][c] > 1 {
				sum += prods[r][c]
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	//part1()
	part2()
}