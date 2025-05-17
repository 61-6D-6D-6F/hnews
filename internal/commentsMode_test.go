package internal

import (
	"testing"
)

// TestCommentsChangeStateNext calls internal.ChangeState with input "n",
// checking for a valid state
func TestCommentsChaToListngeStateNext(t *testing.T) {
	input := "n"
	expCurrentPos := 1

	var state State
	state.CurrentSiblings = []int{0, 1, 2}
	state.CurrentPos = 0

	commentsMode := NewCommentsMode(state)

	newState := commentsMode.ChangeState(input)

	if newState.CurrentPos != expCurrentPos {
		t.Errorf("ChangeState(n) = %d want to match %d", newState.CurrentPos, expCurrentPos)
	}
}

// TestCommnetsChangeStateNextOnLast calls internal.ChangeState with input "n"
// on the last item, checking for a valid state
func TestCommentsChangeStateNextOnLast(t *testing.T) {
	input := "n"
	expCurrentPos := 2

	var state State
	state.CurrentSiblings = []int{0, 1, 2}
	state.CurrentPos = 2

	commentsMode := NewCommentsMode(state)

	newState := commentsMode.ChangeState(input)

	if newState.CurrentPos != expCurrentPos {
		t.Errorf("ChangeState(n) = %d want to match %d", newState.CurrentPos, expCurrentPos)
	}
}

// TestCommentsChangeStatePrev calls internal.ChangeState with input "p",
// checking for a valid state
func TestCommentsChangeStatePrev(t *testing.T) {
	input := "p"
	expCurrentPos := 1

	var state State
	state.CurrentSiblings = []int{0, 1, 2}
	state.CurrentPos = 2

	commentsMode := NewCommentsMode(state)

	newState := commentsMode.ChangeState(input)

	if newState.CurrentPos != expCurrentPos {
		t.Errorf("ChangeState(p) = %d want to match %d", newState.CurrentPos, expCurrentPos)
	}
}

// TestCommentsChangeStatePrevOnFirst calls internal.ChangeState with input "p"
// on the first item, checking for a valid state
func TestCommentsChangeStatePrevOnFirst(t *testing.T) {
	input := "p"
	expCurrentPos := 0

	var state State
	state.CurrentSiblings = []int{0, 1, 2}
	state.CurrentPos = 0

	commentsMode := NewCommentsMode(state)

	newState := commentsMode.ChangeState(input)

	if newState.CurrentPos != expCurrentPos {
		t.Errorf("ChangeState(p) = %d want to match %d", newState.CurrentPos, expCurrentPos)
	}
}

// TestCommentsChangeStateBackToList calls internal.ChangeState with input "b",
// checking for a valid state
func TestCommentsChangeStateBackToList(t *testing.T) {
	input := "b"
	expMode := Details

	var state State
	state.HistorySiblings = [][]int{}

	commentsMode := NewCommentsMode(state)

	newState := commentsMode.ChangeState(input)

	if newState.Mode != expMode {
		t.Errorf("ChangeState(b) = %d want to match %d", newState.Mode, expMode)
	}
}

// TestCommentsChangeStateBackToParent calls internal.ChangeState with input "b",
// checking for a valid state
func TestCommentsChangeStateBackToParent(t *testing.T) {
	input := "b"
	expCurrentSiblings := []int{0, 1}
	expCurrentPos := 1

	var state State
	state.HistorySiblings = [][]int{{0, 1}}
	state.HistoryPos = []int{1}

	commentsMode := NewCommentsMode(state)

	newState := commentsMode.ChangeState(input)

	if len(newState.CurrentSiblings) != len(expCurrentSiblings) ||
		newState.CurrentPos != expCurrentPos {
		t.Errorf("ChangeState(b) = %d, %d want to match %d, %d",
			len(newState.CurrentSiblings), state.CurrentPos,
			len(expCurrentSiblings), expCurrentPos)
	}
}

// TestCommentsChangeStateNoReply calls internal.ChangeState with input "r",
// checking for a valid state
func TestCommentsChangeStateNoReply(t *testing.T) {
	input := "r"
	expCurrentSiblings := []int{0, 1}
	expCurrentPos := 1

	var state State
	state.CurrentSiblings = []int{0, 1}
	state.CurrentPos = 1
	state.FetchedComment.Kids = []int{}

	commentsMode := NewCommentsMode(state)

	newState := commentsMode.ChangeState(input)

	if newState.CurrentPos != expCurrentPos ||
		len(newState.CurrentSiblings) != len(expCurrentSiblings) {
		t.Errorf("ChangeState(r) = %d, %d want to match %d, %d",
			newState.CurrentPos, expCurrentPos,
			len(newState.CurrentSiblings), len(expCurrentSiblings))
	}
}

// TestCommentsChangeStateReply calls internal.ChangeState with input "r",
// checking for a valid state
func TestCommentsChangeStateReply(t *testing.T) {
	input := "r"
	expHistorySiblings := [][]int{{0, 1}}
	expHistoryPos := []int{1}

	var state State
	state.CurrentSiblings = []int{0, 1}
	state.CurrentPos = 1
	state.FetchedComment.Kids = []int{0, 1}

	commentsMode := NewCommentsMode(state)

	newState := commentsMode.ChangeState(input)

	if newState.HistoryPos[0] != expHistoryPos[0] ||
		len(newState.HistorySiblings) != len(expHistorySiblings) {
		t.Errorf("ChangeState(r) = %d, %d want to match %d, %d",
			newState.HistoryPos[0], expHistoryPos[0],
			len(newState.HistorySiblings), len(expHistorySiblings))
	}
}

// TestCommentsChangeStateNotSupported calls internal.ChangeState with input not supported,
// checking for a valid state
func TestCommentsChangeStateNotSupported(t *testing.T) {
	inputs := []string{"i", "3", "44", "6f", "k9"}
	expPageNumber := 22
	expMode := Comments

	var state State
	state.PageNumber = 22
	state.Mode = Comments

	commentsMode := NewCommentsMode(state)

	for _, input := range inputs {
		newState := commentsMode.ChangeState(input)

		if newState.PageNumber != expPageNumber || newState.Mode != expMode {
			t.Errorf("ChangeState(not supported) = %d, %d want to match %d, %d",
				newState.PageNumber, newState.Mode,
				expPageNumber, expMode)
		}
	}
}
