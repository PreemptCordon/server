package discover

import (
	"github.com/preemptcordon/server/obj"
)

var Jobs []obj.DiscoverJob

func init() {
	Jobs = make([]obj.DiscoverJob, 0)

}
func DiscoverQueue() {
	for job := range Jobs {
		isdone := job.evalone()
		if isdone {
			Jobs.pop(job)
		}
		if Jobs.empty {
			break
		}
	}
}
