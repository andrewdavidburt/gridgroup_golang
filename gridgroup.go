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

type Group struct {
	Coords []Coord
}

type Collection struct {
//	Number int
	Grouping []Group
}

// eliminate this global variable
//var group_temp []int
var group_new_temp Group
//var collection_temp Collection

func FloodFillAlgo(grid [][]bool, width int, height int, x int, y int, count int, group Group) ([][]bool, Group) {

	if grid[x][y] == true {
		grid[x][y] = false
//		fmt.Println("==============")
		xy_temp := Coord{X: x, Y: y}
		fmt.Println("GROUP:", count, "ELEMENT:", xy_temp)
		group_new_temp.Coords = append(group_new_temp.Coords, xy_temp)	//[]Coord{xy_temp}
		//fmt.Println("grnew_temp", group_new_temp)
		//group_temp = append(group_temp, count, x, y)
		group.Coords = append(group.Coords, xy_temp)
//		fmt.Println("appending", xy_temp, "to group.Coords:", group.Coords, "of group", group)
//		fmt.Println("==============")
//remove count in flood fill algo later
	if (x > 0) &&		(y > 0) 	{	//NW
						FloodFillAlgo(grid, width, height, x-1, y-1, count, group)}
	if 			(y > 0) 	{	//N
						FloodFillAlgo(grid, width, height, x, y-1, count, group)}
	if (x < width-1) &&	(y > 0) 	{	//NE)
						FloodFillAlgo(grid, width, height, x+1, y-1, count, group)}

	if (x > 0)				{	//W
						FloodFillAlgo(grid, width, height, x-1, y, count, group)}
	if (x < width-1)			{	//E
						FloodFillAlgo(grid, width, height, x+1, y, count, group)}

	if (x > 0) && 		(y < height-1) 	{	//SW
						FloodFillAlgo(grid, width, height, x-1, y+1, count, group)}
	if 			(y < height-1)	{	//S
						FloodFillAlgo(grid, width, height, x, y+1, count, group)}
	if (x < width-1) &&	(y < height-1)	{	//SE
						FloodFillAlgo(grid, width, height, x+1, y+1, count, group)}

        } else {
                return grid, group
        }

//when does this return get called?
	return grid, group
}

func DefineGroups(grid [][]bool) (int) {
	var count int = 0
	var width = 4		//len(grid)
	var height = 4		//len(grid[0])
	fmt.Println("len", len(grid))
	var groupvar []Group
	var group Group
	
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[x][y] == true {
				grid, group = FloodFillAlgo(grid, width, height, x, y, count, group)
				count = count + 1
				fmt.Println("New count:", count)
				fmt.Println("group_new_temp:", group_new_temp)
				fmt.Println("group", group)
				groupvar = append(groupvar, group)
				fmt.Println("groupvar", groupvar)
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
	//fmt.Println(group_temp)
	fmt.Println(group_new_temp)
}
