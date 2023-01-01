package obj

type Option struct {
	decision         Decision
	authorRecommends map[int]Option
	maintainer       EntityObj
	wiki             WikiArticle
}
