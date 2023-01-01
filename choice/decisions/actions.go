package decisions
type ActionPairs struct {
	action Action
	livelyness EvalKind
}
type EvalKind
const (
	Live EvalKind = iota,
	Deadline
)
type Action
type ActionType
const  (
	ManualAction ActionType = iota,
	ChoiceApproveOption,
	ChoiceRejectOption,
	WikiApproveEdit,
	WikiRejectEdit,
	ModDownrankImpose,
	ModDownrankRemove,
	ModHideImpose,
	ModHideRemove,
	ModRemovalImpose,
	ModRemovalRemove,
	ModPreserveImpose,
	ModPreserveRemove,
	ModSilenceNoticesImpose,
	ModSilenceNoticesRemove,
	GroupInviteCreate,
	GroupInviteDelete,
	
)