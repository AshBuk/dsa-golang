// Multiplication Table Generator
// Time: O(n²) - nested loops iterate n×n times
// Space: O(n²) - stores n×n elements in 2D slice

package main

import "fmt"

func multiplicationTable(size int) [][]int {
	table := [][]int{}
	for i := 1; i <= size; i++ {
		row := []int{}
		for j := 1; j <= size; j++ {
			row = append(row, i*j)
		}
		table = append(table, row)
	}
	return table
}

func main() {
	var size int
	fmt.Scan(&size)
	for _, row := range multiplicationTable(size) {
		for _, v := range row {
			fmt.Printf(`%4d`, v)
		}
		fmt.Println()
	}

}

/*
Example result for size = 3:
table = [][]int{
    {1, 2, 3},    // first element table[0] - slice []int{1, 2, 3}
    {2, 4, 6},    // second element table[1] - slice []int{2, 4, 6}
    {3, 6, 9},    // third element table[2] - slice []int{3, 6, 9}
}

How to access elements:
table[0]       // → [1, 2, 3]      (entire first row)
table[1]       // → [2, 4, 6]      (entire second row)
table[2]       // → [3, 6, 9]      (entire third row)

table[0][0]    // → 1              (row 0, column 0)
table[0][1]    // → 2              (row 0, column 1)
table[1][2]    // → 6              (row 1, column 2)
table[2][2]    // → 9              (row 2, column 2)
*/
