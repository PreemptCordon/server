package mydata
type ImportDestination
const (
	Follows ImportDestination = iota,
	Delegates,
	Mutes,
	Blocks,
	Regions,
	Groups,
)

func importcsv(user UserObj, destination ImportDestination, csv string){
	for entry in csv {
		user.lists[destination].queue(entry)
	}
}