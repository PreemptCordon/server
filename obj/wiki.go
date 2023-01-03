package obj

import (
	"github.com/google/uuid"
)

type WikiArticle struct {
	Version           Section
	Key               uuid.UUID
	Category          Category
	Controller        *EntityObj
	ACL               []ACL
	Tags              []Category
	VisitCounter      int       // rolling redis counter of visits, for traffic-control de-ranking or notifications
	HydrateRequest    *Choice   // Decision template to re-hydrate a post
	DehydrateDecision *Decision // how the post dehydrates to S3 out of the DB. usually, based on visit counter being low, or post age
}
type Section struct {
	Title   string
	Summary string
	Body    Markdown
	Owner   *EntityObj
}
type Markdown string
type WikiDB interface {
	Load(id string) WikiArticle
}
type Category struct {
	Owner *EntityObj
}
