package user

import (
	"github.com/preemptcordon/server/obj"
)

type Karma struct {
	user  obj.UserObj
	score int64
}

func KarmaResolve(user obj.UserObj) Karma {
	karma := Karma{user, 10}
	// upvotes = DB.select(user,upvotes)
	// reports = DB.permanent.select(user,reports)
	// content = DB.select(user,content)
	return karma
}
