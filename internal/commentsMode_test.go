package internal

import (
	"testing"
)

func testCommentsChangeState(t *testing.T, input string, states ...[]State) {
	for _, state := range states {
		commentsMode := NewCommentsMode(state[0])

		newState := commentsMode.ChangeState(input)

		if !equalState(newState, state[1]) {
			t.Errorf("ChangeState(%s)\ninitial state:\n%+v\nnew state:\n%+v\nexpected state:\n%+v",
				input, state[0], newState, state[1])
		}
	}
}

// TestCommentsChangeStateNext calls internal.ChangeState with input "n",
// checking for a valid state
func TestCommentsChangeStateNext(t *testing.T) {
	input := "n"

	testCommentsChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// not last comment
		[]State{
			{CurrentPos: 0, CurrentSiblings: []int{0, 1, 2}},
			{CurrentPos: 1, CurrentSiblings: []int{0, 1, 2}}},
		// last comment
		[]State{
			{CurrentPos: 2, CurrentSiblings: []int{0, 1, 2}},
			{CurrentPos: 2, CurrentSiblings: []int{0, 1, 2},
				UiInfo: "Last comment of the comment chain"}},
	)
}

// TestCommentsChangeStatePrev calls internal.ChangeState with input "p",
// checking for a valid state
func TestCommentsChangeStatePrev(t *testing.T) {
	input := "p"

	testCommentsChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// not first comment
		[]State{
			{CurrentPos: 2, CurrentSiblings: []int{0, 1, 2}},
			{CurrentPos: 1, CurrentSiblings: []int{0, 1, 2}}},
		// first comment
		[]State{
			{CurrentPos: 0, CurrentSiblings: []int{0, 1, 2}},
			{CurrentPos: 0, CurrentSiblings: []int{0, 1, 2},
				UiInfo: "First comment of the comment chain"}},
	)
}

// TestCommentsChangeStateBack calls internal.ChangeState with input "b",
// checking for a valid state
func TestCommentsChangeStateBack(t *testing.T) {
	input := "b"

	testCommentsChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// parent is story
		[]State{
			{Mode: Comments, HistorySiblings: [][]int{}},
			{Mode: Details, HistorySiblings: [][]int{}}},
		// parent is comment
		[]State{
			{HistorySiblings: [][]int{{0, 1}}, HistoryPos: []int{1}},
			{CurrentSiblings: []int{0, 1}, CurrentPos: 1}},
	)
}

// TestCommentsChangeStateReply calls internal.ChangeState with input "r",
// checking for a valid state
func TestCommentsChangeStateReply(t *testing.T) {
	input := "r"

	testCommentsChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// no reply
		[]State{
			{CurrentSiblings: []int{0, 1}, CurrentPos: 1,
				FetchedComment: Comment{Kids: []int{}}},
			{CurrentSiblings: []int{0, 1}, CurrentPos: 1,
				UiInfo: "No reply yet"}},

		// has reply
		[]State{
			{CurrentSiblings: []int{0, 1}, CurrentPos: 1,
				HistorySiblings: [][]int{}, HistoryPos: []int{},
				FetchedComment: Comment{Kids: []int{1, 2}}},
			{CurrentSiblings: []int{1, 2}, CurrentPos: 0,
				HistorySiblings: [][]int{{0, 1}}, HistoryPos: []int{1},
				FetchedComment: Comment{Kids: []int{}}}},
	)
}

// TestCommentsChangeStateNotSupported calls internal.ChangeState with input not supported,
// checking for a valid state
func TestCommentsChangeStateNotSupported(t *testing.T) {
	inputs := []string{"i", "3", "44", "6f", "k9"}

	for _, input := range inputs {
		testCommentsChangeState(t, input,
			// [ initial state,
			//   expected state ]

			// not supported
			[]State{
				{Mode: Comments, PageNumber: 2},
				{Mode: Comments, PageNumber: 2,
					UiInfo: "Error: input not supported"}},
		)
	}
}
