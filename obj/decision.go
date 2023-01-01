package obj

import (
	"time"
)

type Decision struct {
	Style           ChoiceType
	Requirements    RequirementsObj
	FacilitatorPool EntityObj // approves edits to voter guide
	DebatorPool     EntityObj // proposes options
	VoterPool       EntityObj // votes count (observers can still react)
	Evaldate        time.Time // ending date
	Options         []Option
	Criteria        map[string]Argument
	Actions         []ActionPair
}
type DecisionResult struct {
	Source    Decision
	Result    string
	Resultref WikiArticle
}
