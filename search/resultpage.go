type ResultStyle
const (
	feed ResultStyle = iota,
	search SearchStyle,
)
type ResultEntity {
	class ObjType
	title string
	key string
	summary string
	wiki string
	score int
}

type ResultPage struct {
	style ResultStyle
	results []ResultEntity
}
func insert(item ResultEntity){
	item.score
}