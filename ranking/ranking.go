type Ranking
const (
	ReverseChron Ranking = iota,
	FeedSort,
	CommentSort,
	DiscoverSort,
)
type Rank interface{
	style Ranking
	Resolve(UserObj) ResultPage
}