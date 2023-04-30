package main

func SpiralTraverse(array [][]int) []int {
	// Initialize an empty slice to hold the result
	result := []int{}
	// If the input array is empty, return the empty result
	if len(array) == 0 {
		return result
	}
	// Initialize variables to keep track of the boundaries of the matrix
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1

	// Traverse the matrix in a spiral order
	for startRow <= endRow && startCol <= endCol {
		// Traverse the top row from left to right
		for col := startCol; col <= endCol; col++ {
			result = append(result, array[startRow][col])
		}
		// Traverse the rightmost column from top to bottom
		for row := startRow + 1; row <= endRow; row++ {
			result = append(result, array[row][endCol])
		}
		// Traverse the bottom row from right to left, if there is more than one row
		for col := endCol - 1; col >= startCol; col-- {
			// If there is only one row left, break the loop to avoid duplicating the elements
			if startRow == endRow {
				break
			}
			result = append(result, array[endRow][col])
		}
		// Traverse the leftmost column from bottom to top, if there is more than one column
		for row := endRow - 1; row > startRow; row-- {
			// If there is only one column left, break the loop to avoid duplicating the elements
			if startCol == endCol {
				break
			}
			result = append(result, array[row][startCol])
		}
		// Update the boundaries of the matrix
		startRow++
		endRow--
		startCol++
		endCol--
	}
	// Return the result slice
	return result
}