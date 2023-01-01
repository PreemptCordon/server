package discover
type DiscoverScale int
const (
	binary DiscoverScale = iota,
	logarithmic,
	linear,
	exponential,
)
type DiscoverFilterDirection
const (
	include DiscoverFilterDirection = iota,
	exclude
)
type DiscoverFilter struct {
	enabled bool,
	scale DiscoverScale,
	zombies bool,
	distance int,
	direction DiscoverFilterDirection,
}
type DiscoverOptions struct {
	disagree_blocked_me DiscoverFilter,
	disagree_blocked_them DiscoverFilter,
	agree_follower DiscoverFilter,
	agree_following DiscoverFilter,
	agree_delegated DiscoverFilter,
	agree_delegating DiscoverFilter,
}
func GetUsers() []UserObj{
	if zombies
}