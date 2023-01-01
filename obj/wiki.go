package obj

type WikiObj struct {
	articles []WikiArticle
}
type WikiArticle struct {
	taxonomy   Category
	controller EntityObj
	acl        []ACL
}
type Section struct {
	text Markdown
}
type Markdown string
