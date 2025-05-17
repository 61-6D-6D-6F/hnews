package internal

import (
	"testing"
)

// TestDetailsChangeStateBack calls internal.ChangeState with input "b",
// checking for a valid state
func TestDetailsChangeStateBack(t *testing.T) {
	input := "b"
	expMode := List

	var state State

	detailsMode := NewDetailsMode(state)

	newState := detailsMode.ChangeState(input)

	if newState.Mode != expMode {
		t.Errorf("ChangeState(b) = %d want to match %d", newState.Mode, expMode)
	}
}

// TestDetailsChangeStateNoComment calls internal.ChangeState with input "c",
// checking for a valid state
func TestDetailsChangeStateNoComment(t *testing.T) {
	input := "c"
	expMode := Details

	var state State
	state.Mode = Details

	detailsMode := NewDetailsMode(state)

	newState := detailsMode.ChangeState(input)

	if newState.Mode != expMode {
		t.Errorf("ChangeState(c) = %d want to match %d", newState.Mode, expMode)
	}
}

// TestDetailsChangeStateComment calls internal.ChangeState with input "c",
// checking for a valid state
func TestDetailsChangeStateComment(t *testing.T) {
	input := "c"
	expCurrentSiblings := []int{0, 1, 2}
	expHistorySiblings := [][]int{}
	expHistoryPos := []int{}
	expCurrentPos := 0
	expMode := Comments

	var state State
	state.SelectedStory.Kids = []int{0, 1, 2}
	state.CurrentSiblings = state.SelectedStory.Kids
	state.HistorySiblings = [][]int{}
	state.HistoryPos = []int{}
	state.CurrentPos = 0
	state.Mode = Comments

	detailsMode := NewDetailsMode(state)

	newState := detailsMode.ChangeState(input)

	if len(newState.CurrentSiblings) != len(expCurrentSiblings) {
		t.Errorf("ChangeState(c) = %d want to match %d",
			len(newState.CurrentSiblings), len(expCurrentSiblings))
	}

	if len(newState.HistorySiblings) != len(expHistorySiblings) {
		t.Errorf("ChangeState(c) = %d want to match %d",
			len(newState.HistorySiblings), len(expHistorySiblings))
	}

	if len(newState.HistoryPos) != len(expHistoryPos) {
		t.Errorf("ChangeState(c) = %d want to match %d",
			len(newState.HistoryPos), len(expHistoryPos))
	}

	if newState.CurrentPos != expCurrentPos {
		t.Errorf("ChangeState(c) = %d want to match %d",
			newState.CurrentPos, expCurrentPos)
	}

	if newState.Mode != expMode {
		t.Errorf("ChangeState(c) = %d want to match %d", newState.Mode, expMode)
	}
}

// TestDetailsChangeStateNotSupported calls internal.ChangeState with input not supported,
// checking for a valid state
func TestDetailsChangeStateNotSupported(t *testing.T) {
	inputs := []string{"i", "3", "44", "6f", "k9"}
	expPageNumber := 22
	expMode := List

	var state State
	state.PageNumber = 22
	state.Mode = List

	detailsMode := NewDetailsMode(state)

	for _, input := range inputs {
		newState := detailsMode.ChangeState(input)

		if newState.PageNumber != expPageNumber || newState.Mode != expMode {
			t.Errorf("ChangeState(not supported) = %d, %d want to match %d, %d",
				newState.PageNumber, newState.Mode,
				expPageNumber, expMode)
		}
	}
}
