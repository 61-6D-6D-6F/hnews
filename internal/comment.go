package internal

type Comment struct {
	By      string `json:"by"`
	Text    string `json:"text"`
	Kids    []int  `json:"kids"`
	Deleted bool   `json:"deleted"`
}
