package main

import (
	"reflect"
	"testing"
)

func TestFillDistances(t *testing.T) {
	rooms := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}
	FillDistances(rooms)
	want := [][]int{
		{3, -1, 0, 1},
		{2, 2, 1, -1},
		{1, -1, 2, -1},
		{0, -1, 3, 4},
	}
	if !reflect.DeepEqual(rooms, want) {
		t.Errorf("rooms = %v, want %v", rooms, want)
	}

	// Unreachable room stays INF; walls stay walls.
	sealed := [][]int{{0, -1, INF}}
	FillDistances(sealed)
	if !reflect.DeepEqual(sealed, [][]int{{0, -1, INF}}) {
		t.Errorf("sealed = %v", sealed)
	}

	// No gates at all: nothing changes.
	noGates := [][]int{{INF, INF}}
	FillDistances(noGates)
	if !reflect.DeepEqual(noGates, [][]int{{INF, INF}}) {
		t.Errorf("no gates = %v", noGates)
	}
}
