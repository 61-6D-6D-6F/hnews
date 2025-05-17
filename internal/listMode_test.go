package internal

import (
	"testing"
)

// TestListChangeStateNext calls internal.ChangeState with input "n",
// checking for a valid state
func TestListChangeStateNext(t *testing.T) {
	input := "n"
	expPageNumber := 1

	var state State

	listMode := NewListMode(state)

	newState := listMode.ChangeState(input)

	if newState.PageNumber != expPageNumber {
		t.Errorf("ChangeState(n) = %d want to match %d", newState.PageNumber, expPageNumber)
	}
}

// TestListChangeStateNextOnLast calls internal.ChangeState with input "n"
// on the last item, checking for a valid state
func TestListChangeStateNextOnLast(t *testing.T) {
	input := "n"
	expPageNumber := MAX_STORIES / NUM_PER_PAGE

	var state State
	state.PageNumber = MAX_STORIES / NUM_PER_PAGE

	listMode := NewListMode(state)

	newState := listMode.ChangeState(input)

	if newState.PageNumber != expPageNumber {
		t.Errorf("ChangeState(n) = %d want to match %d", newState.PageNumber, expPageNumber)
	}
}

// TestListChangeStatePrev calls internal.ChangeState with input "p",
// checking for a valid state
func TestListChangeStatePrev(t *testing.T) {
	input := "p"
	expPageNumber := 2

	var state State
	state.PageNumber = 3

	listMode := NewListMode(state)

	newState := listMode.ChangeState(input)

	if newState.PageNumber != expPageNumber {
		t.Errorf("ChangeState(p) = %d want to match %d", newState.PageNumber, expPageNumber)
	}
}

// TestListChangeStatePrevOnFirst calls internal.ChangeState with input "p"
// on the first item, checking for a valid state
func TestListChangeStatePrevOnFirst(t *testing.T) {
	input := "p"
	expPageNumber := 1

	var state State
	state.PageNumber = 1

	listMode := NewListMode(state)

	newState := listMode.ChangeState(input)

	if newState.PageNumber != expPageNumber {
		t.Errorf("ChangeState(p) = %d want to match %d", newState.PageNumber, expPageNumber)
	}
}

// TestListChangeStateNum calls internal.ChangeState with input "3",
// checking for a valid state
func TestListChangeStateNum(t *testing.T) {
	input := "3"
	expStory := Story{Title: "expPageNumber"}
	expMode := Details

	var state State
	state.FetchedStories = []Story{{}, {}, {Title: "expPageNumber"}}

	listMode := NewListMode(state)

	newState := listMode.ChangeState(input)

	if newState.SelectedStory.Title != expStory.Title || newState.Mode != expMode {
		t.Errorf("ChangeState(3) = %s, %d want to match %s, %d",
			newState.SelectedStory.Title, newState.Mode,
			expStory.Title, expMode)
	}
}

// TestListChangeStateNotSupported calls internal.ChangeState with input not supported,
// checking for a valid state
func TestListChangeStateNotSupported(t *testing.T) {
	inputs := []string{"i", "44", "6f", "k9"}
	expPageNumber := 22
	expMode := List

	var state State
	state.PageNumber = 22
	state.Mode = List

	listMode := NewListMode(state)

	for _, input := range inputs {
		newState := listMode.ChangeState(input)

		if newState.PageNumber != expPageNumber || newState.Mode != expMode {
			t.Errorf("ChangeState(not supported) = %d, %d want to match %d, %d",
				newState.PageNumber, newState.Mode,
				expPageNumber, expMode)
		}
	}
}
