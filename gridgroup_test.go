package main

import (
	"testing"
//	"fmt"
)

// the grid definitions are stored in slices of y- (vertical) stripes
// and are not visual analogs

//TestGrid1 should result in 2 groups
func DefineTestGrid1() [][]bool {
     grid := [][]bool{
             {false, false, false, true},
             {true, false, true, false},
             {true, false, false, false},
             {true, true, false, false},
             }
     return grid
}


//TestGrid2 should result in 3 groups
func DefineTestGrid2() [][]bool {
     grid := [][]bool{
             {true, true, true, true},
             {false, false, false, true},
             {true, true, false, false},
             {true, false, false, true},
             }
     return grid
}

func Test_DefineGrid(t *testing.T) {
	var grid [][]bool = DefineGrid()
	var comparison_grid = [][]bool{
                {true, false, true, true},
                {false, false, true, false},
                {false, true, false, false},
                {true, true, false, false},
                }

	for i, v:= range grid {
		for n, vv:= range v {
        		if vv != comparison_grid[i][n] {
                        	t.Error("Expected:", comparison_grid[i][n], "got:", vv)
                	}
        	}
	}
}

func Test_FloodFillAlgo(t *testing.T) {
	var grid [][]bool = DefineGrid()
	var count int = 0
	var width = 4
	var height = 4
	var group Group

//	var comparison_group Group
//	comparison_group.Coords = append(comparison_group.Coords, Coord{X: 0, Y: 0})
//	comparison_group.Coords = append(comparison_group.Coords, Coord{X: 0, Y: 2})
//	comparison_group.Coords = append(comparison_group.Coords, Coord{X: 1, Y: 2})
//	comparison_group.Coords = append(comparison_group.Coords, Coord{X: 2, Y: 1})
//	comparison_group.Coords = append(comparison_group.Coords, Coord{X: 3, Y: 0})
//	comparison_group.Coords = append(comparison_group.Coords, Coord{X: 3, Y: 1})
//	comparison_group.Coords = append(comparison_group.Coords, Coord{X: 0, Y: 3})

         for x := 0; x < width; x++ {
                 for y := 0; y < height; y++ {
                         if grid[x][y] == true {
			     grid, group = FloodFillAlgo(grid, width, height, x, y, count, group)
			     count = count + 1
	   		     //fmt.Println("grid", grid)
			     //fmt.Println("count", count)
		 	 }
	  	 }
	}
        if count != 2 {
	    t.Error("Expected group count 2, got:", count)
        }

//	fmt.Println("group new temp", group_new_temp)
//	for i:=0; i < len(group_new_temp); i++ {
//                for n, vv:= range v {
//                        if vv != comparison_group[i][n] {
//                                t.Error("Expected:", comparison_group[i][n], "got:", vv)
//                        }
//                }
//        }

}

func Test_DefineGroups(t *testing.T) {
	var grid [][]bool = DefineGrid()
	var groupct int
	groupct = DefineGroups(grid)
	if groupct != 2 {
		t.Error("Expected group count 2, got:", groupct)
	}
}
