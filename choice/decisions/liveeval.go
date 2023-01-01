package decisions

import (
	"github.com/preemptcordon/server/obj"
)

func liveeval(decision obj.Decision) {
	for action := range decision.Actions {
		switch action.kind {
		// mod preemptive actions
		case (obj.ModDownrankImpose):

		case (obj.ModHideImpose):

		case (obj.ModRatelimitImpose):

		}
	}
}
