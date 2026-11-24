package main

import "fmt"

// INF marks an empty room not yet reached from any gate.
const INF = 1 << 30

// FillDistances overwrites every empty room (INF) with its distance to the
// nearest gate (0), leaving walls (-1) alone. Multi-source BFS: all gates
// start in the queue, so the first visit to a room is the shortest route.
func FillDistances(rooms [][]int) {
	type cell struct{ r, c int }
	queue := []cell{}
	for r, row := range rooms {
		for c, v := range row {
			if v == 0 {
				queue = append(queue, cell{r, c})
			}
		}
	}
	dirs := []cell{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			r, c := cur.r+d.r, cur.c+d.c
			if r < 0 || r >= len(rooms) || c < 0 || c >= len(rooms[r]) || rooms[r][c] != INF {
				continue // wall, gate, or already assigned something shorter
			}
			rooms[r][c] = rooms[cur.r][cur.c] + 1
			queue = append(queue, cell{r, c})
		}
	}
}

func main() {
	rooms := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}
	FillDistances(rooms)
	for _, row := range rooms {
		fmt.Println(row)
	}
}
