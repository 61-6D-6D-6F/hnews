package internal

type Story struct {
	Rank    int
	By      string `json:"by"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Text    string `json:"text"`
	Kids    []int  `json:"kids"`
	Deleted bool   `json:"deleted"`
}
