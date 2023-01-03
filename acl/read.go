package acl

import (
	"fmt"

	"github.com/preemptcordon/server/obj"
)

func RenderEntities(user obj.UserObj) []obj.EntityInterface {
	var result []obj.EntityInterface
	return result
}

func CanRead(acl obj.ACL, user obj.UserObj) bool {
	for _, entity := range RenderEntities(user) {
		fmt.Println(entity)
		// if entity.ID == acl.AppliedEntity.ID {
		// 	return acl.AllowView
		// }
	}
	return false
}
