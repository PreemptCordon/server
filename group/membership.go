package group

type membership struct {
	role     RoleObj
	user     UserObj
	decision Decision
	expires  Time
	recall   Choice
}
