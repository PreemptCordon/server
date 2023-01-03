package obj

type Option struct {
	decision         *Decision
	authorRecommends map[int]*Option
	maintainer       EntityInterface
	wiki             *WikiArticle
}
