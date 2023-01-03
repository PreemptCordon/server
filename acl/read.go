package acl

import "github.com/preemptcordon/server/obj"

func RenderEntities(user obj.UserObj) []obj.EntityObj {
	var result []obj.EntityObj
	return result
}

func CanRead(acl obj.ACL, user obj.UserObj) bool {
	for _, entity := range RenderEntities(user) {
		if entity.ID == acl.AppliedEntity.ID {
			return acl.AllowView
		}
	}
	return false
}
