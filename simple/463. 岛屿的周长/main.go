package main

import "fmt"

//func islandPerimeter(grid [][]int) int {
//	res := 0
//	rowMax := len(grid)
//	columnMax := len(grid[0])
//	for rowIndex, row := range grid {
//		for columnIndex, column := range row {
//			if column == 1 {
//				if rowIndex == 0 || grid[rowIndex-1][columnIndex] == 0 {
//					res += 1
//				}
//				if columnIndex == 0 || grid[rowIndex][columnIndex-1] == 0 {
//					res += 1
//				}
//				if rowIndex == rowMax-1 || grid[rowIndex+1][columnIndex] == 0 {
//					res += 1
//				}
//				if columnIndex == columnMax-1 || grid[rowIndex][columnIndex+1] == 0 {
//					res += 1
//				}
//			}
//		}
//	}
//	return res
//}

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				for _, d := range dir4 {
					if x, y := i+d.x, j+d.y; x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
						ans++
					}
				}
			}
		}
	}
	return
}

func main() {
	// [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(islandPerimeter([][]int{{1}}))
}
