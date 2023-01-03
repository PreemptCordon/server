package discover

import (
	"fmt"

	"github.com/preemptcordon/server/obj"
)

var Jobs []obj.DiscoverJob

func init() {
	Jobs = make([]obj.DiscoverJob, 0)

}
func DiscoverQueue() {
	for job := range Jobs {
		fmt.Println(job)
		// isdone := job.Eval()
		// if isdone {
		// 	Jobs.pop(job)
		// }
		// if Jobs.empty {
		// 	break
		// }
	}
}
