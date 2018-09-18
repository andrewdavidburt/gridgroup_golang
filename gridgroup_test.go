package main

import (
	"testing"
//	"fmt"
)

// the grid definitions are stored in slices of y- (vertical) stripes
// and are not visual analogs

//--not implemented yet-- extra TestGrid1 should result in 2 groups
//func DefineTestGrid1() [][]bool {
//     grid := [][]bool{
//             {false, false, false, true},
//             {true, false, true, false},
//             {true, false, false, false},
//             {true, true, false, false},
//             }
//     return grid
//}


//--not implemented yet-- extra TestGrid2 should result in 3 groups
//func DefineTestGrid2() [][]bool {
//     grid := [][]bool{
//             {true, true, true, true},
//             {false, false, false, true},
//             {true, true, false, false},
//             {true, false, false, true},
//             }
//     return grid
//}

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
	group := make([][][]int, 5)

        for x := 0; x < width; x++ {
             for y := 0; y < height; y++ {
                     if grid[x][y] == true {
			     grid, group = FloodFillAlgo(grid, width, height, x, y, count, group)
			     count = count + 1
	        	 }
	  	 }
	}
        if count != 2 {
	    t.Error("Expected group count 2, got:", count)
        }

//	group results test to be added

}

func Test_DefineGroups(t *testing.T) {
	var grid [][]bool = DefineGrid()
	var groupct int
	group := make([][][]int, 5)

//	var comparison_group [5][][]int
//	comparison_group[4][0] = []int{1, 2}

	groupct, group = DefineGroups(grid, group)
	if groupct != 2 {
		t.Error("Expected group count 2, got:", groupct)
	}

//	group results test to be added
//	for i, v:= range group {
//		for ii, vv:= range v {
//			for iii, vvv:= range vv {
//				if vv != comparison_group[i][ii] {
//					t.Error("Expected:", comparison_group[i][ii], "got:", vv)
//				}
//			}
//		}
//	}

}
