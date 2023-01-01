package acl

import (
	_ "github.com/preemptcordon/server/obj"
)

type ACL struct {
	entity  EntityObj
	traffic AutoRateLimit
	limited RateLimit
	view    bool
	edit    bool
	react   bool
}
