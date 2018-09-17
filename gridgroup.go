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
//fmt.Println("Pos:",grid[x][y])
//fmt.Println("NW")
	if (x > 0) &&		(y > 0) 	{FloodFillAlgo(grid, width, height, x-1, y-1)} //NW
//fmt.Println("N")
	if 			(y > 0) 	{FloodFillAlgo(grid, width, height, x, y-1)} //N
//fmt.Println("NE")
	if (x < width-1) &&	(y > 0) 	{FloodFillAlgo(grid, width, height, x+1, y-1)} //NE

//fmt.Println("W")
	if (x > 0)				{FloodFillAlgo(grid, width, height, x-1, y)} //W
//fmt.Println("E")
	if (x < width-1)			{FloodFillAlgo(grid, width, height, x+1, y)} //E

//fmt.Println("SW")
	if (x > 0) && 		(y < height-1) 	{FloodFillAlgo(grid, width, height, x-1, y+1)} //SW
//fmt.Println("S")
	if 			(y < height-1)	{FloodFillAlgo(grid, width, height, x, y+1)} //S
//fmt.Println("SE")
	if (x < width-1) &&	(y < height-1)	{FloodFillAlgo(grid, width, height, x+1, y+1)} //SE

        } else {
//fmt.Println("returngrid")
                return grid
        }
//this seems unnecessary but not sure how to avoid it
//fmt.Println("weirdreturngrid")
	return grid
}

func DefineGroups(grid [][]bool) (int, [][]bool) {
	var count int = 0
	var width = len(grid)-1
//fmt.Println("width:", width)
	var height = len(grid[0])-1
//fmt.Println("height:", height)
//	var newGrid = [][]bool
	
fmt.Println("Start DefineGroups")
	for x := 0; x <= width; x++ {
//fmt.Println("inside x", x)
		for y := 0; y <= height; y++ {
//fmt.Println("inside y", y)
			if grid[x][y] == true {
//fmt.Println("grid_define_groups:", grid)
				grid = FloodFillAlgo(grid, width, height, x, y)
				count = count + 1
fmt.Println("New count:", count)
fmt.Println("Grid:", grid)
			}
		}
	}

	return count, grid
}

func main() {
	var grid [][]bool = DefineGrid()
fmt.Println("Grid:", grid)
	var groupct int
	var groupdef [][]bool
	//change this and in definegroups output to [][]int to define groups?
	groupct, groupdef = DefineGroups(grid)
	fmt.Println("Group Count:", groupct)
	fmt.Println("Group Definitions:", groupdef)
}
