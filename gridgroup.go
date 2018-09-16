//grid adjacent+diagonal grouping program using floodfill algorithm
package main

import (
	"fmt"
)

//func GenerateGrid() {
//}

//func CountMultipleGrids() {
//}

//func DisplayVisualGridGroups(grid [][]bool) {
//	fmt.Println(grid)
//         for x := range width {
//                 for y := range height {
//                         if grid[x][y] == true {
//				
//                 }
//        }
//}

func DefineGrid() [][]bool {
	grid := [][]bool{
		{1, 0, 0, 1},
		{0, 0, 1, 1},
		{1, 1, 0, 0},
		{1, 0, 0, 0}
		}
	return grid
}

func FloodFillAlgo(grid [][]bool, width int, height int) [][]bool {
	if grid[x][y] == true {
		grid[x][y] = false
	} else {
		break
	}

//double-check these logics
	if (x > 0) &&		(y > 0) 	{FloodFillAlgo(grid, x-1, y-1)} //NW
	if 			(y > 0) 	{FloodFillAlgo(grid, x, y-1)} //N
	if (x < width-1) &&	(y > 0) 	{FloodFillAlgo(grid, x+1, y-1)} //NE

	if (x > 0)				{FloodFillAlgo(grid, x-1, y)} //W
	if (x < width-1)			{FloodFillAlgo(grid, x+1, y)} //E

	if (x > 0) && 		(y < height-1) 	{FloodFillAlgo(grid, x-1, y+1)} //SW
	if 			(y < height-1)	{FloodFillAlgo(grid, x, y+1)} //S
	if (x < width-1) &&	(y < height-1)	{FloodFillAlgo(grid, x+1, y+1)} //SE

}

func DefineGroups(grid [][]bool) int [][]int {
	var count int = 0
	var width = len(grid)
	var height = len(grid[0]
//	var newGrid = [][]bool
	
	for x := range width {
		for y := range height {
			if grid[x][y] == true {
				grid = FloodFillAlgo(grid, width, height)
				count := count + 1
			}
		}
	}

	return count
}

func main() {
	var grid [][]bool = DefineGrid()
//	var groupcount int, groupdef [][]int = DefineGroups(grid)
	var groupcount int = DefineGroups(grid)
	fmt.Println("Group Count:", groupct)
//	fmt.Println("Group Definitions:", groupdef
}

//func name(in int), []out {
//	var output []int
//	for i := 0; i < 10; i++ {
//		if i%2 == 0 {
//			fmt.Println(i, "is even.")
//			output = append(output, i)
//		}
//	}
//	return output
//}
//func main () {
//	var evens []int
//	evens = name(10)
//	fmt.Println("All evens:", evens)
