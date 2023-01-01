package obj

import (
	"time"
)

type Decision struct {
	style        ChoiceType
	requirements RequirementsObj
	facilitator  EntityObj // approves edits to voter guide
	debator      EntityObj // proposes options
	voter        EntityObj // votes count (observers can still react)
	evaldate     time.Time // ending date
	options      []Option
	criteria     map[string]Argument
	actions      []ActionPair
}
type DecisionResult struct {
	source    Decision
	result    string
	resultref WikiArticle
}
