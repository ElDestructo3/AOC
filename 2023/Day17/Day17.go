package main

import (
	"bufio"
	"fmt"
	"os"
	"container/heap"
)

func twotoone(x, y int) int {
	return (x) * 10000 + (y)
}

func onetotwo(n int) (int, int) {
	return n / 10000, n % 10000
}


type Item struct {
	value    [3]int 
	priority int    
	index int 
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  
	item.index = -1 
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value [3]int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

	

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := make([]int, 0)
		for i := 0; i < len(text); i++ {
			arr = append(arr, int(text[i]) - int('0'))
		}
		grid = append(grid, arr)
	}
	start_x := 0
	start_y := 0

	pq := make(PriorityQueue, 0)

	heap.Init(&pq)
	heap.Push(&pq, &Item{value: [3]int{twotoone(start_x, start_y), 0, -1}, priority: 0})


	dist_map := make(map[[3]int]int)
	dist_map[[3]int{twotoone(start_x, start_y), 0, -1}] = 0

	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		pos := item.value[0]
		same_dir_steps := item.value[1]
		dir := item.value[2]
		dist := item.priority
		if dist_map[[3]int{pos, same_dir_steps, dir}] != dist {
			continue
		}
		x, y := onetotwo(pos)
		if x == len(grid) - 1 && y == len(grid[0]) - 1 {
			fmt.Println(dist)
			break
		}
		
		for i := 0; i < len(dirs); i++ {
			dx := x + dirs[i][0]
			dy := y + dirs[i][1]
			if (dir == 0 && i == 1) || (dir == 1 && i == 0) || (dir == 2 && i == 3) || (dir == 3 && i == 2) {
				continue
			}
			if dx < 0 || dx >= len(grid) || dy < 0 || dy >= len(grid[0]) {
				continue
			}
			if i == dir {
				if same_dir_steps == 3 {
					continue
				}
				s, ok := dist_map[[3]int{twotoone(dx, dy), same_dir_steps + 1, i}]
				if ok && s <= dist + grid[dx][dy]{
					continue
				}
				heap.Push(&pq, &Item{value: [3]int{twotoone(dx, dy), same_dir_steps + 1, i}, priority: dist + grid[dx][dy]})
				dist_map[[3]int{twotoone(dx, dy), same_dir_steps + 1, i}] = dist + grid[dx][dy]
				continue
			} 
			s, ok := dist_map[[3]int{twotoone(dx, dy), 1, i}]
			if ok && s <= dist + grid[dx][dy] {
				continue
			}
			heap.Push(&pq, &Item{value: [3]int{twotoone(dx, dy), 1, i}, priority: dist + grid[dx][dy]})
			dist_map[[3]int{twotoone(dx, dy), 1, i}] = dist + grid[dx][dy]
		}
	}

}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := make([]int, 0)
		for i := 0; i < len(text); i++ {
			arr = append(arr, int(text[i]) - int('0'))
		}
		grid = append(grid, arr)
	}
	start_x := 0
	start_y := 0

	pq := make(PriorityQueue, 0)

	heap.Init(&pq)
	heap.Push(&pq, &Item{value: [3]int{twotoone(start_x, start_y), 0, -1}, priority: 0})


	dist_map := make(map[[3]int]int)
	dist_map[[3]int{twotoone(start_x, start_y), 0, -1}] = 0

	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		pos := item.value[0]
		same_dir_steps := item.value[1]
		dir := item.value[2]
		dist := item.priority
		if dist_map[[3]int{pos, same_dir_steps, dir}] != dist {
			continue
		}
		x, y := onetotwo(pos)
		if x == len(grid) - 1 && y == len(grid[0]) - 1 && same_dir_steps >= 4{
			fmt.Println(dist)
			break
		}
		
		for i := 0; i < len(dirs); i++ {
			dx := x + dirs[i][0]
			dy := y + dirs[i][1]
			if (dir == 0 && i == 1) || (dir == 1 && i == 0) || (dir == 2 && i == 3) || (dir == 3 && i == 2) {
				continue
			}
			if dx < 0 || dx >= len(grid) || dy < 0 || dy >= len(grid[0]) {
				continue
			}

			if i == dir {
				if same_dir_steps >= 10 {
					continue
				}
				s, ok := dist_map[[3]int{twotoone(dx, dy), same_dir_steps + 1, i}]
				if ok && s <= dist + grid[dx][dy]{
					continue
				}
				heap.Push(&pq, &Item{value: [3]int{twotoone(dx, dy), same_dir_steps + 1, i}, priority: dist + grid[dx][dy]})
				dist_map[[3]int{twotoone(dx, dy), same_dir_steps + 1, i}] = dist + grid[dx][dy]
				continue
			} 
			if same_dir_steps < 4 && dir != -1{
				continue
			}
			s, ok := dist_map[[3]int{twotoone(dx, dy), 1, i}]
			if ok && s <= dist + grid[dx][dy] {
				continue
			}
			heap.Push(&pq, &Item{value: [3]int{twotoone(dx, dy), 1, i}, priority: dist + grid[dx][dy]})
			dist_map[[3]int{twotoone(dx, dy), 1, i}] = dist + grid[dx][dy]
		}
	}


}

func main() {
	//part1()
	part2()
}