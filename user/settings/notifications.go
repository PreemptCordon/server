package settings

import (
	_ "github.com/preemptcordon/server/notifications"
)

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
