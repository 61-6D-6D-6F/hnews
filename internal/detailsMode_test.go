package internal

import (
	"testing"
)

func testDetailsChangeState(t *testing.T, input string, testCases ...[]State) {
	for _, testCase := range testCases {
		detailsMode := NewDetailsMode(testCase[0])

		newState := detailsMode.ChangeState(input)

		if !equalState(newState, testCase[1]) {
			t.Errorf("ChangeState(%s)\ninitial state:\n%+v\nnew state:\n%+v\nexpected state:\n%+v",
				input, testCase[0], newState, testCase[1])
		}
	}
}

// TestDetailsChangeStateBack calls internal.ChangeState with input "b",
// checking for a valid state
func TestDetailsChangeStateBack(t *testing.T) {
	input := "b"

	testDetailsChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// return to list
		[]State{
			{Mode: Details},
			{Mode: List}},
	)
}

// TestDetailsChangeStateComment calls internal.ChangeState with input "c",
// checking for a valid state
func TestDetailsChangeStateComment(t *testing.T) {
	input := "c"

	testDetailsChangeState(t, input,
		// [ initial state,
		//   expected state ]

		// no comment
		[]State{
			{Mode: Details, SelectedStory: Story{Kids: []int{}}},
			{Mode: Details, UiInfo: "No comment yet"}},
		// has comment
		[]State{
			{SelectedStory: Story{Kids: []int{0, 1, 2}},
				CurrentSiblings: []int{},
				HistorySiblings: [][]int{},
				HistoryPos:      []int{},
				CurrentPos:      0,
				Mode:            Details},
			{SelectedStory: Story{Kids: []int{0, 1, 2}},
				CurrentSiblings: []int{0, 1, 2},
				HistorySiblings: [][]int{},
				HistoryPos:      []int{},
				CurrentPos:      0,
				Mode:            Comments}},
	)
}

// TestDetailsChangeStateNotSupported calls internal.ChangeState with input not supported,
// checking for a valid state
func TestDetailsChangeStateNotSupported(t *testing.T) {
	inputs := []string{"i", "44", "6f", "k9", "n2", "3 e", "k 7"}

	for _, input := range inputs {
		testDetailsChangeState(t, input,
			// [ initial state,
			//   expected state ]

			// not supported
			[]State{
				{Mode: Details, PageNumber: 2},
				{Mode: Details, PageNumber: 2, UiInfo: "Error: input not supported"}},
		)
	}
}
