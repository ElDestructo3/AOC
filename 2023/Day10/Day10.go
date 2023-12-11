package main

import (
	"bufio"
	"fmt"
	"os"
)

func twotoone(x, y int) int {
	return (x + 1) * 10000 + (y + 1)
}

func onetotwo(n int) (int, int) {
	return n / 10000 - 1, n % 10000 - 1
}

//func 


func part1() (map[int]rune, map[int]int, map[int]int) {
	scanner := bufio.NewScanner(os.Stdin)
	start_x := -1
	start_y := -1
	// map from int to int
	dist := make(map[int]int)
	mp := make(map[int]rune)
	parents := make(map[int]int)
	row := 0
	//row_len := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		//row_len = len(text)
		for i, c := range text {
			if c == '.' {
				n := twotoone(row, i)
				mp[n] = c
				dist[n] = 1000000000
				continue
			} else if c == 'S' {
				start_x = row
				start_y = i
				n := twotoone(row, i)
				mp[n] = c
			} else {
				n := twotoone(row, i)
				mp[n] = c
				dist[n] = 1000000000
			}
		}
		row++
	}
	//fmt.Println(start_x, start_y)
	dist[twotoone(start_x, start_y)] = 0
	parents[twotoone(start_x, start_y)] = -1

	queue := make([]int, 0)
	queue = append(queue, twotoone(start_x, start_y))
	var ans int

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		x, y := onetotwo(curr)
		//fmt.Println(x, y)
		char := mp[curr]
		if char == '|' {
			if mp[twotoone(x - 1, y)] == 'S' || mp[twotoone(x + 1, y)] == 'S' {
				if dist[twotoone(x,y)] > 1 {
					ans = dist[twotoone(x,y)]
					break
				}
			}
			par_x, _ := onetotwo(parents[curr])
			if par_x == x - 1 {
				if mp[twotoone(x + 1, y)] != '.' {
					dist[twotoone(x + 1, y)] = dist[curr] + 1
					parents[twotoone(x + 1, y)] = curr
					queue = append(queue, twotoone(x + 1, y))
				}
			} else if par_x == x + 1 {
				if mp[twotoone(x - 1, y)] != '.' {
					dist[twotoone(x - 1, y)] = dist[curr] + 1
					parents[twotoone(x - 1, y)] = curr
					queue = append(queue, twotoone(x - 1, y))
				}
			} else {
				continue
			}
		} else if char == '-' {
			if mp[twotoone(x, y - 1)] == 'S' || mp[twotoone(x, y + 1)] == 'S' {
				if dist[twotoone(x,y)] > 1 {
					ans = dist[twotoone(x,y)]
					break
				}
			}
			_, par_y := onetotwo(parents[curr])
			if par_y == y - 1 {
				if mp[twotoone(x, y + 1)] != '.' {
					dist[twotoone(x, y + 1)] = dist[curr] + 1
					parents[twotoone(x, y + 1)] = curr
					queue = append(queue, twotoone(x, y + 1))
				}
			} else if par_y == y + 1 {
				if mp[twotoone(x, y - 1)] != '.' {
					dist[twotoone(x, y - 1)] = dist[curr] + 1
					parents[twotoone(x, y - 1)] = curr
					queue = append(queue, twotoone(x, y - 1))
				}
			} else {
				continue
			}

		} else if char == '7' {
			if mp[twotoone(x, y - 1)] == 'S' || mp[twotoone(x + 1, y)] == 'S' {
				if dist[twotoone(x,y)] > 1 {
					ans = dist[twotoone(x,y)]
					break

				}
			}
			par_x, par_y := onetotwo(parents[curr])
			if par_x == x + 1 {
				if mp[twotoone(x, y - 1)] != '.' {
					dist[twotoone(x, y - 1)] = dist[curr] + 1
					parents[twotoone(x, y - 1)] = curr
					queue = append(queue, twotoone(x, y - 1))
				}
			} else if par_y == y - 1 {
				if mp[twotoone(x + 1, y)] != '.' {
					dist[twotoone(x + 1, y)] = dist[curr] + 1
					parents[twotoone(x + 1, y)] = curr
					queue = append(queue, twotoone(x + 1, y))
				}
			} else {
				continue
			}

		} else if char == 'L' {
			if mp[twotoone(x, y + 1)] == 'S' || mp[twotoone(x - 1, y)] == 'S' {
				if dist[twotoone(x,y)] > 1 {
					ans = dist[twotoone(x,y)]
					break
				}
			}
			par_x, par_y := onetotwo(parents[curr])
			if par_x == x - 1 {
				if mp[twotoone(x, y + 1)] != '.' {
					dist[twotoone(x, y + 1)] = dist[curr] + 1
					parents[twotoone(x, y + 1)] = curr
					queue = append(queue, twotoone(x, y + 1))
				}
			} else if par_y == y + 1 {
				if mp[twotoone(x - 1, y)] != '.' {
					dist[twotoone(x - 1, y)] = dist[curr] + 1
					parents[twotoone(x - 1, y)] = curr
					queue = append(queue, twotoone(x - 1, y))
				}
			} else {
				continue
			}
		} else if char == 'J' {
			if mp[twotoone(x, y - 1)] == 'S' || mp[twotoone(x - 1, y)] == 'S' {
				if dist[twotoone(x,y)] > 1 {
					ans = dist[twotoone(x,y)]
					break
				}
			}
			par_x, par_y := onetotwo(parents[curr])
			if par_x == x - 1 {
				if mp[twotoone(x, y - 1)] != '.' {
					dist[twotoone(x, y - 1)] = dist[curr] + 1
					parents[twotoone(x, y - 1)] = curr
					queue = append(queue, twotoone(x, y - 1))
				}
			}
			if par_y == y - 1 {
				if mp[twotoone(x - 1, y)] != '.' {
					dist[twotoone(x - 1, y)] = dist[curr] + 1
					parents[twotoone(x - 1, y)] = curr
					queue = append(queue, twotoone(x - 1, y))
				}
			} else {
				continue
			}
		} else if char == 'F' {
			if mp[twotoone(x, y + 1)] == 'S' || mp[twotoone(x + 1, y)] == 'S' {
				if dist[twotoone(x,y)] > 1 {
					ans = dist[twotoone(x,y)]
					break
				}
			}
			par_x, par_y := onetotwo(parents[curr])
			if par_x == x + 1 {
				if mp[twotoone(x, y + 1)] != '.' {
					dist[twotoone(x, y + 1)] = dist[curr] + 1
					parents[twotoone(x, y + 1)] = curr
					queue = append(queue, twotoone(x, y + 1))
				}
			} else if par_y == y + 1 {
				if mp[twotoone(x + 1, y)] != '.' {
					dist[twotoone(x + 1, y)] = dist[curr] + 1
					parents[twotoone(x + 1, y)] = curr
					queue = append(queue, twotoone(x + 1, y))
				}
			} else {
				continue
			}


		} else if char == 'S' {
			// enqueue x, y + 1 and x + 1, y
			if mp[twotoone(x, y + 1)] != '.' && dist[twotoone(x, y + 1)] > dist[curr] + 1 {
				dist[twotoone(x, y + 1)] = dist[curr] + 1
				parents[twotoone(x, y + 1)] = curr
				queue = append(queue, twotoone(x, y + 1))
			}
			
		}
	}
	
	if ans % 2 == 1 {
		fmt.Println(ans / 2 + 1)
	} else {
		fmt.Println(ans / 2)
	}
	return mp, dist, parents
	
}

func part2() {
	//mp, dist, parents := part1()
	mp, dist, _ := part1()
	// bounding_rec_x := 1000000000
	// bounding_rec_y := 1000000000
	// bounding_rec_x2 := -2
	// bounding_rec_y2 := -2

	len_row := 0
	num_row := 0
	for k, _ := range mp {
		x, y := onetotwo(k)
		if y > len_row {
			len_row = y
		}
		if x > num_row {
			num_row = x
		}
		// }
		// if dist[twotoone(x, y)] == 1000000000 {
		// 	continue
		// }
		// //fmt.Println(x, y)
		
		// if x < bounding_rec_x {
		// 	bounding_rec_x = x
		// }
		// if y < bounding_rec_y {
		// 	bounding_rec_y = y
		// }
		// if x > bounding_rec_x2 {
		// 	bounding_rec_x2 = x
		// }
		// if y > bounding_rec_y2 {
		// 	bounding_rec_y2 = y
		// }
	}
	// bounding_rec_x--
	// bounding_rec_y--
	// bounding_rec_x2++
	// bounding_rec_y2++

	// fmt.Println(bounding_rec_x, bounding_rec_y, bounding_rec_x2, bounding_rec_y2)
	// //check if each point is in the bounding rectangle and inside the maze
	// count := 0
	// for i := bounding_rec_x + 1; i < bounding_rec_x2; i++ {
	// 	for j := bounding_rec_y + 1; j < bounding_rec_y2; j++ {
	// 		if dist[twotoone(i, j)] != 1000000000 {
	// 			//fmt.Println(dist[twotoone(i, j)])
	// 			continue
	// 		}
	// 		//fmt.Println(i, j)
	// 		visited := make(map[int]bool)
	// 		unvisited := make([]int, 0)
	// 		unvisited = append(unvisited, twotoone(i, j))
	// 		unvisited_map := make(map[int]bool)
	// 		unvisited_map[twotoone(i, j)] = true
	// 		dirs := [][]int {{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	// 		fmt.Println("Starting BFS")
	// 		for {
	// 			if len(unvisited) == 0 {
	// 				fmt.Print("Found: ")
	// 				fmt.Println(i, j)
	// 				count++
	// 				break
	// 			}
	// 			curr := unvisited[0]
	// 			unvisited = unvisited[1:]
	// 			visited[curr] = true
	// 			x, y := onetotwo(curr)
	// 			fmt.Println(x, y)
	// 			if x == bounding_rec_x || x == bounding_rec_x2 || y == bounding_rec_y || y == bounding_rec_y2 {
	// 				break
	// 			}
	// 			for _, dir := range dirs {
	// 				new_x := x + dir[0]
	// 				new_y := y + dir[1]
	// 				if dist[twotoone(new_x, new_y)] == 1000000000 && !visited[twotoone(new_x, new_y)] && !unvisited_map[twotoone(new_x, new_y)] {
	// 					unvisited = append(unvisited, twotoone(new_x, new_y))
	// 					unvisited_map[twotoone(new_x, new_y)] = true
	// 				}
	// 				if _, ok := mp[twotoone(new_x, new_y)]; !ok {
	// 					unvisited = append(unvisited, twotoone(new_x, new_y))
	// 					unvisited_map[twotoone(new_x, new_y)] = true
	// 				}
	// 			}
	// 		}
	// 		// if len(unvisited) == 0 {
	// 		// 	count++
	// 		// 	fmt.Println(i, j)
	// 		// }
	// 	}
	// }
	// fmt.Println(count)
	count := 0
	for x := 0; x <= num_row; x++ {
		for y := 0; y <= len_row; y++ {
			if dist[twotoone(x, y)] != 1000000000 {
				continue
			}
			crosses := 0
			x2, y2 := x, y
			for x2 <= num_row && y2 <= len_row {
				char := mp[twotoone(x2, y2)]
				if dist[twotoone(x2, y2)] != 1000000000 && char != 'L' && char != '7'{
					crosses++
				}
				x2++
				y2++
			}
			if crosses % 2 == 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	//part1()
	part2()
}