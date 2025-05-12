package stories

type Story struct {
	Rank        int
	By          string `json:"by"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Text        string `json:"text"`
	Kids        []int  `json:"kids"`
	Descendants int    `json:"descendants"`
}

type StoryMode int

const (
	List StoryMode = iota
	Detail
)

var Mode = List

var CurrentPage = 1
