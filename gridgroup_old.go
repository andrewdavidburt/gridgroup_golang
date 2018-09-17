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

type Coord struct {
	X, Y int
}

type Group struct {
	Coords []Coord
}

func FloodFillAlgo(grid [][]bool, width int, height int, x int, y int, groupvar Group) ([][]bool, Group) {
	if grid[x][y] == true {
		grid[x][y] = false
		//add to a struct or array to return
//		var loc Coord
//		loc.X = x
//		loc.Y = y
//		var loc Coord
//		loc = Coord{X: x, Y: y}
//		fmt.Println(loc)
////		loc = append(loc, {X: x, Y: y})

		coordvar := Coord{X: x, Y: y}
//		coordslicevar := []Coord{}
//		groupvar := Group{coordslicevar}
		groupvar.Coords = append(groupvar.Coords, coordvar)
		
		//this needs to be sent back up through all recursive calls

//fmt.Println("Pos:", x, y)

	//fmt.Println("NW")
	if (x > 0) &&		(y > 0) 	{FloodFillAlgo(grid, width, height, x-1, y-1, groupvar)} //NW
	//fmt.Println("N")
	if 			(y > 0) 	{FloodFillAlgo(grid, width, height, x, y-1, groupvar)} //N
	//fmt.Println("NE")
	if (x < width-1) &&	(y > 0) 	{FloodFillAlgo(grid, width, height, x+1, y-1, groupvar)} //NE

	//fmt.Println("W")
	if (x > 0)				{FloodFillAlgo(grid, width, height, x-1, y, groupvar)} //W
	//fmt.Println("E")
	if (x < width-1)			{FloodFillAlgo(grid, width, height, x+1, y, groupvar)} //E

	//fmt.Println("SW")
	if (x > 0) && 		(y < height-1) 	{FloodFillAlgo(grid, width, height, x-1, y+1, groupvar)} //SW
	//fmt.Println("S")
	if 			(y < height-1)	{FloodFillAlgo(grid, width, height, x, y+1, groupvar)} //S
	//fmt.Println("SE")
	if (x < width-1) &&	(y < height-1)	{FloodFillAlgo(grid, width, height, x+1, y+1, groupvar)} //SE

        } else {

fmt.Println("1returngrid:", x, y, groupvar)
//fmt.Println("grid:", grid)
                return grid, groupvar
        }
//when does this return get used??
fmt.Println("2returngrid:", x, y, groupvar)
//fmt.Println("grid:", grid)
	return grid, groupvar
}

func DefineGroups(grid [][]bool, coordslicevar []Coord, groupvar Group) (int, Group) {
	var count int = 0
	var width = len(grid)-1
//fmt.Println("width:", width)
	var height = len(grid[0])-1
//fmt.Println("height:", height)
//	groupdef:= make([]GroupCoord)
//	var groupdef []int
//        var loc []Group
//	coordslicevar := []Coord{}
//	groupvar := Group{coordslicevar}
//	tempgroupvar := groupvar

	
//fmt.Println("Start DefineGroups")
	for x := 0; x <= width; x++ {
//fmt.Println("inside x", x)
		for y := 0; y <= height; y++ {
//fmt.Println("inside y", y)
			if grid[x][y] == true {
//fmt.Println("grid_define_groups:", grid)
				fmt.Println("before:", groupvar)
				grid, groupvar = FloodFillAlgo(grid, width, height, x, y, groupvar)
				fmt.Println("after:", groupvar)
				count = count + 1
//				groupdef = append(groupdef, count)
				
fmt.Println("New count:", count)
fmt.Println("Grid:", grid)
			}
		}
	}

	return count, groupvar
}

func main() {
	var grid [][]bool = DefineGrid()
fmt.Println("Grid:", grid)
	var groupct int
//	var groupdef []int
	coordslicevar := []Coord{}
	groupvar := Group{coordslicevar}
//	groupct, groupdef = DefineGroups(grid)
	groupct, groupvar = DefineGroups(grid, coordslicevar, groupvar)
	fmt.Println("FINAL RESULTS")
	fmt.Println("Group Count:", groupct)
	fmt.Println("Group Definitions:", groupvar)
}
