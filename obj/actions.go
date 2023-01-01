package obj

type ActionPair struct {
	action     Action
	livelyness EvalKind
}
type EvalKind int

const (
	Live EvalKind = iota
	Deadline
)

type Action int
type ActionType int

const (
	ManualAction ActionType = iota
	ChoiceApproveOption
	ChoiceRejectOption
	WikiApproveEdit
	WikiRejectEdit
	ModDownrankImpose
	ModDownrankRemove
	ModHideImpose
	ModHideRemove
	ModRemovalImpose
	ModRemovalRemove
	ModPreserveImpose
	ModPreserveRemove
	ModSilenceNoticesImpose
	ModSilenceNoticesRemove
	GroupInviteCreate
	GroupInviteDelete
)
