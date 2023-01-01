package obj

type RankingPref int

const (
	ReverseChron RankingPref = iota
	FeedSort
	CommentSort
	DiscoverSort
)

type Rank interface {
	Resolve(UserObj) ResultPage
}
type RankWeight int64

const (
	LinearWeight RankWeight = iota
	BinaryWeight
	LogarithmicWeight
	ExponentialWeight
)
