package mydata

import "github.com/preemptcordon/server/obj"

type ImportDestination int

const (
	Follows ImportDestination = iota
	Delegates
	Mutes
	Blocks
	Regions
	Groups
)

func importcsv(user obj.UserObj, destination ImportDestination, csv string) {
	for entry := range csv {
		user.lists[destination].queue(entry)
	}
}
