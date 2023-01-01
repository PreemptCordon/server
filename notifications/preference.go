package notifications

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
