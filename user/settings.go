package user

import (
	_ "../notifications"
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
	ranking   UserRankDefaulPreferences
	notices   UserNoticeDefaultPreferences
	discovery DiscoverPreferences
}
