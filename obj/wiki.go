package obj

import (
	"github.com/google/uuid"
)

type WikiArticle struct {
	Version    Section
	Key        uuid.UUID
	Category   Category
	Controller *EntityObj
	ACL        []ACL
	Tags       []Category
}
type Section struct {
	Text Markdown
}
type Markdown string
type WikiDB interface {
	Load(id string) WikiArticle
}
type Category struct {
	owner *EntityObj
}
