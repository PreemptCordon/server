package obj

type EntityObj struct {
	handle  string
	objtype ObjType
	about   WikiObj
}
type EntityInterface interface {
	Follow(EntityObj, EntityObj) bool
	Unfollow(EntityObj, EntityObj) bool
	Notifications() []string
	Queue(EntityObj)
}
type UserObj struct {
	Name   string
	ID     string
	Lists  UserDataLists
	Email  string
	Handle string
}

func (u UserObj) Notifications() []string {
	result := []string{
		"test",
		"one",
		"two",
		"three",
		"four",
	}
	return result
}

type EntityAction struct {
	list []EntityObj
}

type UserDataLists struct {
	Follows   []EntityObj
	Delegates []EntityObj
	Blocks    []EntityObj
	Mutes     []EntityObj
}
type UserDataListsInterface interface {
	Queue(EntityObj)
}

type profile struct {
	user UserObj
	wiki WikiObj
}
type UserList []string
type UserListsData struct {
	MuteList     UserList
	BlockList    UserList
	FollowList   UserList
	DelegateList UserList
}
type RankPreference string
type UserRankDefaultPreferences struct {
	FeedSort         RankPreference
	CommentSort      RankPreference
	SearchSort       RankPreference
	DiscoverSort     RankPreference
	ReverseChronSort RankPreference
}

type UserSettings struct {
	mydata    UserListsData
	ranking   UserRankDefaultPreferences
	notices   UserNoticeDefaultPreferences
	discovery DiscoverPreferences
}
type DiscoverPreferences struct {
	discoverable bool `false`
	traversable  bool `false`
}

type UserNoticeDefaultPreferences struct {
	// queues
	DiscoverResult     NoticePreference
	DeletionResult     NoticePreference
	TransparencyReport NoticePreference
	// social actions
	FollowRequest NoticePreference
	NewFollower   NoticePreference
	NewDelegator  NoticePreference
	Invites       NoticePreference
	NewMembership NoticePreference
	// choices
	NewChoice      NoticePreference
	NewOption      NoticePreference
	OptionStatus   NoticePreference // approved, rejected
	NewArgument    NoticePreference
	ArgumentStatus NoticePreference // approved, rejected
	NewDecision    NoticePreference

	// mod transparency
	SearchDownrank NoticePreference
	SearchRemoved  NoticePreference
	// report transparency
	ReportNotice NoticePreference
	ReportAppeal NoticePreference
	ReportResult NoticePreference
}
type NoticePreference int64

const (
	Disabled NoticePreference = iota
	Weekly
	Daily
	Instant
)

type NoticeSetting struct {
	BrowserNotify NoticePreference
	EmailNotify   NoticePreference
}
