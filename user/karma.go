package user

type Karma struct {
	user UserObject
	score int64
}

func KarmaResolve(user UserObject) Karma {
	karma := Karma{user}
	upvotes = DB.select(user,upvotes)
	reports = DB.permanent.select(user,reports)
	content = DB.select(user,content)
	return karma
}
