package obj

type EntityObj struct {
	handle  string
	objtype ObjType
	about   WikiObj
}
type EntityInterface interface {
	follow(EntityObj, EntityObj) bool
	unfollow(EntityObj, EntityObj) bool
}
