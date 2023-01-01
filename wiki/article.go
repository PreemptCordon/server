package wiki

import (
	"net/http"

	_ "github.com/preemptcordon/server/acl"
)

func articlehandler(w http.ResponseWriter, r *http.Request) {
	check := acl.Lookup(r.URL.Path)
	return check
}
