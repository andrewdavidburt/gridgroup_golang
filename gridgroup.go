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
		{true, false, false, true},
		{false, false, true, true},
		{true, true, false, false},
		{true, false, false, false},
		}
	return grid
}

func FloodFillAlgo(grid [][]bool, width int, height int, x int, y int) [][]bool {
	if grid[x][y] == true {
		grid[x][y] = false

	if (x > 0) &&		(y > 0) 	{FloodFillAlgo(grid, width, height, x-1, y-1)} //NW
	if 			(y > 0) 	{FloodFillAlgo(grid, width, height, x, y-1)} //N
	if (x < width-1) &&	(y > 0) 	{FloodFillAlgo(grid, width, height, x+1, y-1)} //NE

	if (x > 0)				{FloodFillAlgo(grid, width, height, x-1, y)} //W
	if (x < width-1)			{FloodFillAlgo(grid, width, height, x+1, y)} //E

	if (x > 0) && 		(y < height-1) 	{FloodFillAlgo(grid, width, height, x-1, y+1)} //SW
	if 			(y < height-1)	{FloodFillAlgo(grid, width, height, x, y+1)} //S
	if (x < width-1) &&	(y < height-1)	{FloodFillAlgo(grid, width, height, x+1, y+1)} //SE

        } else {
                return grid
        }
//this seems unnecessary but not sure how to avoid it
	return grid
}

func DefineGroups(grid [][]bool) (int, [][]bool) {
	var count int = 0
	var width = len(grid)
	var height = len(grid[0])
//	var newGrid = [][]bool
	
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			if grid[x][y] == true {
				grid = FloodFillAlgo(grid, width, height, x, y)
				count = count + 1
			}
		}
	}

	return count, grid
}

func main() {
	var grid [][]bool = DefineGrid()
	var groupct int
	var groupdef [][]bool
	//change this and in definegroups output to [][]int to define groups?
	groupct, groupdef = DefineGroups(grid)
	fmt.Println("Group Count:", groupct)
	fmt.Println("Group Definitions:", groupdef)
}
