package decisions
/* decisions are individually rendered choices */

struct Decision {
	style ChoiceType
	requirements RequirementsObj
	facilitator EntityObj // approves edits to voter guide
	debator EntityObj // proposes options
	voter EntityObj // votes count (observers can still react)
	evaldate Date // ending date
	options []Options
	criteria string[Arguments]
	actions []ActionPairs
}
struct DecisionResult {
	source Decision
	result string
	resultref Wiki
}