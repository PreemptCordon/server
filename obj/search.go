package obj

type ResultStyle int

const (
	FeedResultStyle ResultStyle = iota
	SearchResultStyle
)

type ResultEntity struct {
	Class   ObjType
	Title   string
	Key     string
	Summary string
	Wiki    string
	Score   int
}

type ResultPage struct {
	Style   ResultStyle
	Results []ResultEntity
}
