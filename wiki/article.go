package wiki

import (
	"net/http"

	_ "github.com/preemptcordon/server/acl"
)

func articlehandler(w http.ResponseWriter, r *http.Request) {
	check := acl.lookup(r.URL.Path)
	return check
}
