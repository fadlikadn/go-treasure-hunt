package main

import (
	"fmt"
)

/**
# -> Obstacle
. -> Clear Path
X -> Initial Player starting positing
E -> Improvement, target location
$ -> possible treasure location
*/

const R = 6
const C = 8

var matrix = [R][C]string{
	{"#", "#", "#", "#", "#", "#", "#", "#"},
	{"#", ".", ".", ".", ".", ".", ".", "#"},
	{"#", ".", "#", "#", "#", ".", ".", "#"},
	{"#", ".", ".", ".", "#", ".", "#", "#"},
	{"#", "X", "#", ".", ".", ".", "E", "#"},
	{"#", "#", "#", "#", "#", "#", "#", "#"},
}
var possibility = [R][C]string{
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
	{".", ".", ".", ".", ".", ".", ".", "."},
}
var r, c int
var sr, sc int
var i, j int
var rQueue []int
var cQueue []int
var moveCount int
var nodesLeftInLayer int
var nodesInNextLayer int
var reachedEnd bool
var (
	intVal = map[bool]int{true: 1}
)

var visited = [R][C]bool{
	{false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false},
}

// direction: up, right, bottom, left
var dr = []int{-1, 0, +1, 0}
var dc = []int{0, +1, 0, -1}
var move = []int{0, 0, 0, 0}

func main() {
	for i = 0; i < R; i++ {
		for j = 0; j < C; j++ {
			if matrix[i][j] == "X" {
				sr = i
				sc = j
				break
			}
		}
	}

	moveCount = 0
	nodesLeftInLayer = 1
	nodesInNextLayer = 0

	reachedEnd = false

	solved := solve()
	fmt.Printf("Number of step: %d\n", solved)
	fmt.Println("Direction Statistic :")
	fmt.Printf("Top: %d, Right: %d, Bottom: %d, Left: %d", move[0], move[1], move[2], move[3])
	fmt.Println()
	fmt.Println()

	showPossibleTreasureLocation()
}

func showPossibleTreasureLocation() {
	var possibleTreasureHunt []map[string]int
	for i = 0; i < R; i++ {
		for j = 0; j < C; j++ {
			if matrix[i][j] == "." || matrix[i][j] == "X" || matrix[i][j] == "E" {
				possible := map[string]int{
					"rows":    i,
					"columns": j,
				}
				possibility[i][j] = "$"
				possibleTreasureHunt = append(possibleTreasureHunt, possible)
			}
		}
	}

	fmt.Println("Probable coordinate points where treasure might be located:")
	for _, e := range possibleTreasureHunt {
		fmt.Printf("(%d,%d) ", e["rows"], e["columns"])
	}
	fmt.Println()
	fmt.Println()

	for i = 0; i < R; i++ {
		for j = 0; j < C; j++ {
			fmt.Printf(possibility[i][j])
		}
		fmt.Println()
	}
}

func solve() int {
	rQueue = append(rQueue, sr)
	cQueue = append(cQueue, sc)
	visited[sr][sc] = true
	showVisitedTable()
	for len(rQueue) > 0 {
		r = rQueue[0]
		c = cQueue[0]
		rQueue = rQueue[1:]
		cQueue = cQueue[1:]
		if matrix[r][c] == "E" {
			reachedEnd = true
			break
		}
		exploreNeighbour(r, c)
		nodesLeftInLayer--
		if nodesLeftInLayer == 0 {
			nodesLeftInLayer = nodesInNextLayer
			nodesInNextLayer = 0
			moveCount++
		}
	}
	if reachedEnd {
		return moveCount
	}
	return -1
}

func showVisitedTable() {
	for i = 0; i < R; i++ {
		for j = 0; j < C; j++ {
			fmt.Printf("%d", intVal[visited[i][j]])
		}
		fmt.Println()
	}
	fmt.Println()
}

func exploreNeighbour(r int, c int) {
	for i = 0; i < 4; i++ {
		rr := r + dr[i]
		cc := c + dc[i]

		if rr < 0 || cc < 0 {
			continue
		}
		if rr >= R || cc >= C {
			continue
		}

		if visited[rr][cc] {
			continue
		}
		if matrix[rr][cc] == "#" {
			continue
		}

		rQueue = append(rQueue, rr)
		cQueue = append(cQueue, cc)
		visited[rr][cc] = true
		move[i]++
		showVisitedTable()
		nodesInNextLayer++
	}
}
