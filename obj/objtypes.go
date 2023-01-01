package obj

type ObjType int64

const (
	WikiType ObjType = iota
	SectionType
	UserType
	GroupType
	RegionType
	ChoiceType
	DecisionType
	RequirementsType
	OptionType
	CategoryType
)

type Category struct {
	owner EntityObj
}
