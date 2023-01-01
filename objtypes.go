type ObjType int64
const (
	Wiki ObjType = iota,
	Section,
	User,
	Group,
	Region,
	Choice,
	Decision,
	Requirements,
	Option,
	
)
type EntityObj struct {
	handle string,
	objtype ObjType,
	about Wiki,
}
type EntityType interface {
	follow(EntityObj, EntityType) bool
	unfollow(EntityObj, EntityType) bool
}
const (
	UserType EntityObj = iota,
	GroupType,
	CategoryType,
)