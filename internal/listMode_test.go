package internal

import (
	"testing"
)

func testListChangeState(t *testing.T, input string, states ...[]State) {
	for _, state := range states {
		listMode := NewListMode(state[0])

		newState := listMode.ChangeState(input)

		if !equalState(newState, state[1]) {
			t.Errorf("ChangeState(%s)\ninitial state:\n%+v\nnew state:\n%+v\nexpected state:\n%+v",
				input, state[0], newState, state[1])
		}
	}
}

// TestListChangeStateNext calls internal.ChangeState with input "n",
// checking for a valid state
func TestListChangeStateNext(t *testing.T) {
	input := "n"

	testListChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// not last page
		[]State{
			{PageNumber: 1},
			{PageNumber: 2}},
		// last page
		[]State{
			{PageNumber: MAX_STORIES / NUM_PER_PAGE},
			{PageNumber: MAX_STORIES / NUM_PER_PAGE}},
	)
}

// TestListChangeStatePrev calls internal.ChangeState with input "p",
// checking for a valid state
func TestListChangeStatePrev(t *testing.T) {
	input := "p"

	testListChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// not first page
		[]State{
			{PageNumber: 2},
			{PageNumber: 1}},
		// first page
		[]State{
			{PageNumber: 1},
			{PageNumber: 1}},
	)
}

// TestListChangeStateNum calls internal.ChangeState with input "2",
// checking for a valid state
func TestListChangeStateNum(t *testing.T) {
	input := "2"

	testListChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// number 1-9
		[]State{
			{Mode: List, SelectedStory: Story{},
				FetchedStories: []Story{{}, {Title: "exp"}}},
			{Mode: Details, SelectedStory: Story{Title: "exp"},
				FetchedStories: []Story{{}, {Title: "exp"}}}},
	)
}

// TestListChangeStateNotSupported calls internal.ChangeState with input not supported,
// checking for a valid state
func TestListChangeStateNotSupported(t *testing.T) {
	inputs := []string{"i", "44", "6f", "k9"}

	for _, input := range inputs {
		testListChangeState(t, input,
			// [ initial state,
			//   expected state ]

			// not supported
			[]State{
				{Mode: List, PageNumber: 2},
				{Mode: List, PageNumber: 2}},
		)
	}
}
