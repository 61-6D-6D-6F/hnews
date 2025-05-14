package internal

type Comment struct {
	By      string `json:"by"`
	Text    string `json:"text"`
	Kids    []int  `json:"kids"`
	Parent  int    `json:"parent"`
	Type    string `json:"type"`
	Deleted bool   `json:"deleted"`
}
