package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func shoelace(vertices [][2]int) int {  //determinant of triangle formula extended to n vertices
	sum := 0
	for i := 0; i < len(vertices) - 1; i++ {
		sum += vertices[i][0] * vertices[i + 1][1]
		sum -= vertices[i][1] * vertices[i + 1][0]
	}
	sum += vertices[len(vertices) - 1][0] * vertices[0][1]
	sum -= vertices[len(vertices) - 1][1] * vertices[0][0]
	if sum < 0 {
		sum *= -1
	}
	sum /= 2
	return sum
}

func picks_theorem_lattice_points(vertices [][2]int, perimeter int) int {
	// states that area = (perimeter) /2 - 1 + (number of lattice points inside polygon)
	area := shoelace(vertices)
	return (perimeter / 2) + 1 + area
	
}

	

func part1() {
	scanner := bufio.NewScanner(os.Stdin)
	start_x := 0
	start_y := 0
	vertices := make([][2]int, 0)
	perimeter := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := strings.Split(text, " ")
		edge_len, _ := strconv.Atoi(arr[1])
		if arr[0] == "R" {
			start_y += edge_len
		} else if arr[0] == "L" {
			start_y -= edge_len
		} else if arr[0] == "U" {
			start_x -= edge_len
		} else if arr[0] == "D" {
			start_x += edge_len
		}
		perimeter += edge_len
		vertices = append(vertices, [2]int{start_x, start_y})
	}
	fmt.Println(picks_theorem_lattice_points(vertices, perimeter))
	

}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)
	start_x := 0
	start_y := 0
	vertices := make([][2]int, 0)
	perimeter := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		arr := strings.Split(text, " ")
		var code rune
		if arr[2][len(arr[2]) - 2] == '0' {
			code = 'R'
		} else if arr[2][len(arr[2]) - 2] == '2' {
			code = 'L'
		} else if arr[2][len(arr[2]) - 2] == '3' {
			code = 'U'
		} else if arr[2][len(arr[2]) - 2] == '1' {
			code = 'D'
		}
		edge_len := 0
		for i := 0; i < len(arr[2]) - 4; i++ {
			edge_len *= 16
			dig_val := 0
			if arr[2][i + 2] >= 'a' && arr[2][i + 2] <= 'f' {
				dig_val = int(arr[2][i + 2]) - int('a') + 10
			} else {
				dig_val = int(arr[2][i + 2]) - int('0')
			}
			edge_len += dig_val
		}
		if code == 'R' {
			start_y += edge_len
		} else if code == 'L' {
			start_y -= edge_len
		} else if code == 'U' {
			start_x -= edge_len
		} else if code == 'D' {
			start_x += edge_len
		}

		perimeter += edge_len
		vertices = append(vertices, [2]int{start_x, start_y})
	}
	fmt.Println(picks_theorem_lattice_points(vertices, perimeter))
	


}

func main() {
	//part1()
	part2()
}