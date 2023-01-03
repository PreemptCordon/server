package obj

import "github.com/google/uuid"

type RoleEntityObj struct {
	handle string
	group  *GroupObj
	ACL    []ACL
}

// table: role:user (2 way relation)
// table: ACL:role (2 way relation)
// table: ACL:read,write,vote,discover
// table: user:subject:email(daily,instant,weekly,none),browser(instantly,daily,weekly,none) <- mail pref resolved from most specific to least specific (decision, group, tags, default)
// table: blocklist: blocker:blockee (2 way)
// mute list (one way): feed user:ignored user: expire time : reason
// downrank list: article, factor, decision, notes(section), author(user)

type GroupObj struct {
	handle string
	about  *WikiArticle
	ID     uuid.UUID
}

// 	Groups can have regions, or might not.

type EntityInterface interface {
	Follow(EntityInterface, EntityInterface) bool
	Unfollow(EntityInterface, EntityInterface) bool
	Notifications() []string
	Queue(EntityInterface)
}
type UserObj struct {
	Name     string
	ID       string
	Lists    UserDataLists
	Email    string
	Handle   string
	Settings UserSettings
	Summary  *Section
	Wiki     *WikiArticle
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
	list []EntityInterface
}

type UserDataLists struct {
	Follows   []EntityInterface
	Delegates []EntityInterface
	Blocks    []EntityInterface
	Mutes     []EntityInterface
}
type UserDataListsInterface interface {
	Queue(EntityInterface)
}

type UserList []string
type UserListsData struct {
	MuteList     UserList
	BlockList    UserList
	FollowList   UserList
	DelegateList UserList
}
type UserRankDefaultPreferences struct {
	FeedSort         RankingPref
	CommentSort      RankingPref
	SearchSort       RankingPref
	DiscoverSort     RankingPref
	ReverseChronSort RankingPref
}

type UserSettings struct {
	MyData    UserListsData
	Ranking   UserRankDefaultPreferences
	Notices   UserNoticeDefaultPreferences
	Discovery DiscoverPreferences
}
type DiscoverPreferences struct {
	Discoverable bool `false`
	Traversable  bool `false`
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
type DiscoverScale int

const (
	binary DiscoverScale = iota
	logarithmic
	linear
	exponential
)

type DiscoverFilterDirection bool

const (
	include DiscoverFilterDirection = true
	exclude DiscoverFilterDirection = false
)

type DiscoverFilter struct {
	enabled   bool
	scale     DiscoverScale
	zombies   bool
	distance  int
	direction DiscoverFilterDirection
}
type DiscoverOptions struct {
	disagree_blocked_me   DiscoverFilter
	disagree_blocked_them DiscoverFilter
	agree_follower        DiscoverFilter
	agree_following       DiscoverFilter
	agree_delegated       DiscoverFilter
	agree_delegating      DiscoverFilter
}
