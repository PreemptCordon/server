package wiki

import (
	"net/http"

	_ "github.com/preemptcordon/server/acl"
)
type 

type WikiArticle struct {
	taxonomy   Category
	controller EntityObj
	acl []ACL
}

func articlehandler(w http.ResponseWriter, r *http.Request) {
	check := acl.lookup(r.URL.Path)
	return check
}
