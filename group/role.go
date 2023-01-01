package group

type RoleObj struct {
	handle  string
	group   GroupObj
	term    timeduration
	duties  Wiki
	recall  Choice
	picking Choice
}
