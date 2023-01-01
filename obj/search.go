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
type UnrankedResult ResultEntity
type SearchConsider struct {
	title   bool
	summary bool
	author  bool
	bio     bool
	groups  bool
}
type TopicLimit struct {
	topic    Category
	distance int
}
type SearchTerms struct {
	query     string
	consider  SearchConsider
	distlimit []TopicLimit
}
