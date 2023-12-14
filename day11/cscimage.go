package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/iyassou/advent-of-code-2023/internal"
)

type CSCImage struct {
	NumRows     int
	RowIndex    []int
	ColumnIndex []int
}

func NewCSCImage(input []byte) (*CSCImage, error) {
	// ASSUMPTION: input is rectangular.
	lines := internal.Lines(input)
	m, n := len(lines), len(lines[0])
	c := &CSCImage{
		NumRows:     m,
		RowIndex:    []int{},
		ColumnIndex: []int{},
	}
	total, previous, runLength := 0, 0, 1
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			pixel := lines[row][col]
			if pixel == '#' {
				total++
				c.RowIndex = append(c.RowIndex, row)
			}
		}
		if total == previous {
			runLength++
		} else {
			c.ColumnIndex = append(c.ColumnIndex, runLength, previous)
			previous = total
			runLength = 1
		}
		if col == n-1 {
			c.ColumnIndex = append(c.ColumnIndex, runLength, previous)
		}
	}
	return c, nil
}

func (c *CSCImage) NumColumns() int {
	if c == nil {
		return 0
	}
	sum := 0
	for i := 2; i < len(c.ColumnIndex); i += 2 {
		sum += c.ColumnIndex[i]
	}
	return sum
}

func (c *CSCImage) galaxyCount() int {
	if c == nil {
		return 0
	}
	return c.ColumnIndex[len(c.ColumnIndex)-1]
}

func (c *CSCImage) dimensions() (int, int) {
	if c == nil {
		return 0, 0
	}
	return c.NumRows, c.NumColumns()
}

func (c *CSCImage) getRowIndex(column int) int {
	if c == nil || !(0 <= column && column <= c.NumColumns()) {
		return -1
	}
	colSum := 0
	for i := 0; i < len(c.ColumnIndex); i += 2 {
		runLength := c.ColumnIndex[i]
		if colSum <= column && column < colSum+runLength {
			return c.ColumnIndex[i+1]
		}
		colSum += runLength
	}
	return c.galaxyCount() + 1
}

func (c *CSCImage) galaxyAt(x, y int) (bool, error) {
	if c == nil {
		return false, errors.New("nil CSCImage")
	}
	if !(0 <= x && x < c.NumRows) {
		return false, fmt.Errorf("invalid x coordinate %d", x)
	}
	if !(0 <= y && y < c.NumColumns()) {
		return false, fmt.Errorf("invalid y coordinate %d", y)
	}
	min, max := c.getRowIndex(y), c.getRowIndex(y+1)
	if min < 0 || max < 0 {
		return false, fmt.Errorf("unexpected error occurred (min=%d, max=%d)", min, max)
	}
	for _, v := range c.RowIndex[min:max] {
		if v == x {
			return true, nil
		}
	}
	return false, nil
}

func (c *CSCImage) String() string {
	if c == nil {
		return ""
	}
	var sb strings.Builder
	M, N := c.dimensions()
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if ok, err := c.galaxyAt(i, j); err != nil {
				sb.WriteByte('?')
			} else if ok {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (c *CSCImage) emptyRows() []int {
	if c == nil {
		return nil
	}
	notEmpty := map[int]bool{}
	for _, row := range c.RowIndex {
		notEmpty[row] = true
	}
	empty := make([]int, c.NumRows-len(notEmpty))
	i := 0
	for row := 0; row < c.NumRows; row++ {
		if _, contains := notEmpty[row]; !contains {
			empty[i] = row
			i++
		}
	}
	return empty
}

func (c *CSCImage) emptyColumns() []int {
	if c == nil {
		return nil
	}
	empty := []int{}
	colSum := 0
	for i := 0; i < len(c.ColumnIndex); i += 2 {
		runLength := c.ColumnIndex[i]
		for j := colSum + 1; j < colSum+runLength; j++ {
			empty = append(empty, j-1)
		}
		colSum += runLength
	}
	return empty
}

func (c *CSCImage) expandEmptyRows(by int) {
	if c == nil {
		return
	}
	rows := c.emptyRows()
	if rows == nil {
		return
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i] < rows[j] })
	insertionIndex := func(x int) int {
		for i, row := range rows {
			if row > x {
				return i
			}
		}
		return len(rows)
	}
	for i, row := range c.RowIndex {
		if row <= rows[0] {
			continue
		}
		c.RowIndex[i] += insertionIndex(row) * by
	}
	c.NumRows += len(rows) * by
}

func (c *CSCImage) expandEmptyColumns(by int) {
	if c == nil {
		return
	}
	for i := 0; i < len(c.ColumnIndex); i += 2 {
		runLength := c.ColumnIndex[i]
		if runLength > 1 {
			c.ColumnIndex[i] += by * (runLength - 1)
		}
	}
}

func (c *CSCImage) ExpandUniverse(by int) {
	c.expandEmptyRows(by)
	c.expandEmptyColumns(by)
}

func (c *CSCImage) GalaxyLocations() [][2]int {
	if c == nil {
		return nil
	}
	g, loc := 0, make([][2]int, c.galaxyCount())
	N := c.NumColumns()
	for col := 0; col < N; col++ {
		min, max := c.getRowIndex(col), c.getRowIndex(col+1)
		for _, row := range c.RowIndex[min:max] {
			loc[g] = [2]int{row, col}
			g++
		}
	}
	return loc
}
