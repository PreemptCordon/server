package obj

type ResultStyle int

const (
	FeedResultStyle ResultStyle = iota
	SearchResultStyle
)

type Rankable interface {
	RankSelf(UserObj, Rankable) Rankable
}
type ResultEntity struct {
	// Class   ObjType / could also have meant kind
	Title   string
	Key     string
	Summary string
	Wiki    string
	Score   int
	Author  EntityInterface
	ACL     ACL
}

type ResultPage struct {
	Style   ResultStyle
	Results []ResultEntity
}
type UnrankedResult ResultEntity
type SearchConsider struct {
	Title   bool
	Summary bool
	Author  bool
	Bio     bool
	Groups  bool
}
type TopicLimit struct {
	Topic    Category
	Distance int
}
type SearchTerms struct {
	Query          string
	Consider       SearchConsider
	DistTopicLimit []TopicLimit
	Tags           []Category
}
