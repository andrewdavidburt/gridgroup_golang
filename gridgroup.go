//grid adjacent+diagonal grouping program using floodfill algorithm
package main

import (
	"fmt"
)

//func DefineGrid() [][]bool {
//	grid := [][]bool{
//		{true, false, false, true},
//		{false, false, true, true},
//		{true, true, false, false},
//		{true, false, false, false},
//		}
//	return grid
//}

func DefineGrid() [][]bool {			 // swapped x/y in grid
        grid := [][]bool{
                {true, false, true, true},
                {false, false, true, false},
                {false, true, false, false},
                {true, true, false, false},
                }
        return grid
}

type Coord struct {
	X, Y int
}

type Group struct {
	Coords []Coord
}

var group_temp []int

func FloodFillAlgo(grid [][]bool, width int, height int, x int, y int, count int) ([][]bool) {

	if grid[x][y] == true {
		//fmt.Println("______________")
		//fmt.Println(grid)
		//fmt.Println(grid[x][y])	
		grid[x][y] = false
		//fmt.Println(grid[x][y])
		//fmt.Println("ff 1 tf change start:", x, y)
		//fmt.Println("ff 1 tf change stgrd:", grid)
		fmt.Println("==============")
		fmt.Println("GROUP:", count, "ELEMENT:", x, ",", y)
		group_temp = append(group_temp, count, x, y)
		fmt.Println("==============")

	if (x > 0) &&		(y > 0) 	{//fmt.Println("NW from", x, y, grid[x][y])
						FloodFillAlgo(grid, width, height, x-1, y-1, count)}
	if 			(y > 0) 	{//fmt.Println("N from ", x, y, grid[x][y])
						FloodFillAlgo(grid, width, height, x, y-1, count)}
	if (x < width-1) &&	(y > 0) 	{//fmt.Println("NE from", x, y, grid[x][y])
						FloodFillAlgo(grid, width, height, x+1, y-1, count)}

	if (x > 0)				{//fmt.Println("W from", x, y, grid[x][y])
						FloodFillAlgo(grid, width, height, x-1, y, count)}
	if (x < width-1)			{//fmt.Println("E from", x, y, grid[x][y])
						FloodFillAlgo(grid, width, height, x+1, y, count)}

	if (x > 0) && 		(y < height-1) 	{//fmt.Println("SW from", x, y, grid[x][y])
						FloodFillAlgo(grid, width, height, x-1, y+1, count)}
	if 			(y < height-1)	{//fmt.Println("S from", x, y, grid[x][y])
						FloodFillAlgo(grid, width, height, x, y+1, count)}
	if (x < width-1) &&	(y < height-1)	{//fmt.Println("SE from", x, y, grid[x][y])
						FloodFillAlgo(grid, width, height, x+1, y+1, count)}

        } else {
		//fmt.Println("ff 1 returngrid:", x, y, grid)
                return grid
        }

	//when does this return get called?
	//fmt.Println("ff 2 ff nochange start:")
	//fmt.Println("ff 2 returngrid:", x, y, grid)
	return grid
}

func DefineGroups(grid [][]bool) (int) {
	var count int = 0
	var width = 4//len(grid)-1
	var height = 4//len(grid[0])-1
	
	for x := 0; x < width; x++ {
	//fmt.Println("outer x loop", x)
		for y := 0; y < height; y++ {
		//fmt.Println("outer y loop", y, "within loop x", x)
			if grid[x][y] == true {
				//fmt.Println("outer true:", x, y)
				grid = FloodFillAlgo(grid, width, height, x, y, count)
				count = count + 1
				fmt.Println("New count:", count)
				//fmt.Println("Grid:", grid)
			}
		}
	}

	return count
}

func main() {
	var grid [][]bool = DefineGrid()
	fmt.Println("Grid:", grid)
	var groupct int
	groupct = DefineGroups(grid)
	fmt.Println("FINAL RESULTS")
	fmt.Println("Group Count:", groupct)
	fmt.Println(group_temp)
}
