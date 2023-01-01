package user

import (
	_ "github.com/preemptcordon/server/user/settings"
)

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
