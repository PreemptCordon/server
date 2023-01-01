package user
type EntityAction struct {
	list []EntityObj
}
type Interface EntityAction {
	queue(EntityObj)
}

type UserDataLists struct {
	Follows   []EntityObj
	Delegates []EntityObj
	Blocks    []EntityObj
	Mutes     []EntityObj
}
type Interface UserDataLists struct {
	queue(EntityObj)
}

type UserObject struct {
	Name  string
	ID    string
	lists UserDataLists
}
