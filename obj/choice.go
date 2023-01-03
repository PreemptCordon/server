package obj

type LazySelection struct {
	user      UserObj
	option    Option
	magnitude int
}
type LazyArg int

const (
	AgreeWillHelp LazyArg = iota
	AgreeWontHelp
	DisagreeNoAlt
	DisagreeAlt
)

type ChoiceType int64

const (
	CompromiseRCV ChoiceType = iota
	ConsensusUnanimus
	LazyConsensus
	ThresholdVote
)

type ChoiceDecide interface {
	Create()
	Eval()
}
type Choice struct {
	requirements *RequirementsObj
	template     WikiArticle
}
