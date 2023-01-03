package obj

import (
	"time"
)

type Decision struct {
	Style           ChoiceType
	Requirements    RequirementsObj
	FacilitatorPool EntityInterface // approves edits to voter guide
	DebatorPool     EntityInterface // proposes options
	VoterPool       EntityInterface // votes count (observers can still react)
	Evaldate        time.Time       // ending date
	Options         []Option
	Criteria        map[string]Argument
	Actions         []ActionPair
}
type DecisionResult struct {
	Source    Decision
	Result    string
	Resultref WikiArticle
}
