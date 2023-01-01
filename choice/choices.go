type ChoiceType int64
const (
	CompromiseRCV ChoiceType iota,
	ConsensusUnanimus ChoiceType,
	LazyConsensus ChoiceType,
	ThresholdVote ChoiceType
)
type ChoiceDecide interface {
	Create()
	Eval()
}

