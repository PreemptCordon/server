package wiki

import (
	"net/http"

	"github.com/preemptcordon/server/acl"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	check := acl.Lookup(r.URL.Path)
	if check == false {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("forbidden"))
	}
	return
}
