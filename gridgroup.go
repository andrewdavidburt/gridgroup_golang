//grid adjacent+diagonal grouping program using an implementation of floodfill algorithm
package main

import (
	"fmt"
)

// the grid definition is stored in slices of y- (vertical) stripes
// and are not visual analogs
// in this case, it results in the same visual grid as you used:
// X O O X
// X O X X
// X X O O
// X O O O
// and should also result in 2 groupings of adjacent tiles

func DefineGrid() [][]bool {
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

//type Group struct {
//	Coords []Coord
//}

//type Collection struct {
//	Grouping []Group
//}

func FloodFillAlgo(grid [][]bool, width int, height int, x int, y int, count int, group_3 [][]int) ([][]bool, [][]int) {

	if grid[x][y] == true {
		grid[x][y] = false
		xy_temp := Coord{X: x, Y: y}
		fmt.Println("GROUP:", count, "ELEMENT:", xy_temp)
		//group_new_temp.Coords = append(group_new_temp.Coords, xy_temp)	//[]Coord{xy_temp}

		group_3[count] = append(group_3[count], x, y)
		
	if (x > 0) &&		(y > 0) 	{	//NW
						FloodFillAlgo(grid, width, height, x-1, y-1, count, group_3)}
	if 			(y > 0) 	{	//N
						FloodFillAlgo(grid, width, height, x, y-1, count, group_3)}
	if (x < width-1) &&	(y > 0) 	{	//NE)
						FloodFillAlgo(grid, width, height, x+1, y-1, count, group_3)}

	if (x > 0)				{	//W
						FloodFillAlgo(grid, width, height, x-1, y, count, group_3)}
	if (x < width-1)			{	//E
						FloodFillAlgo(grid, width, height, x+1, y, count, group_3)}

	if (x > 0) && 		(y < height-1) 	{	//SW
						FloodFillAlgo(grid, width, height, x-1, y+1, count, group_3)}
	if 			(y < height-1)	{	//S
						FloodFillAlgo(grid, width, height, x, y+1, count, group_3)}
	if (x < width-1) &&	(y < height-1)	{	//SE
						FloodFillAlgo(grid, width, height, x+1, y+1, count, group_3)}

        } else {
                return grid, group_3
        }

	return grid, group_3
}

func DefineGroups(grid [][]bool, group_3 [][]int) (int, [][]int) {
	var count int = 0
	var width = 4		//len(grid)
	var height = 4		//len(grid[0])

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[x][y] == true {
				grid, group_3 = FloodFillAlgo(grid, width, height, x, y, count, group_3)
				count = count + 1
			}
		}
	}

	return count, group_3
}

func main() {
	var grid [][]bool = DefineGrid()
	group_3 := make([][]int, 5)
	var groupct int
	groupct, group_3 = DefineGroups(grid, group_3)
	fmt.Println("output group_3", group_3)
	fmt.Println("Group Count:", groupct)
}
