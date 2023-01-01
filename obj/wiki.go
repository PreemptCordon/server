package obj

import (
	"github.com/google/uuid"
)

type WikiObj struct {
	articles []WikiArticle
	Version  Section
	Key      uuid.UUID
}
type WikiArticle struct {
	taxonomy   Category
	controller EntityObj
	acl        []ACL
}
type Section struct {
	Text Markdown
}
type Markdown string
