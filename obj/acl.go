package obj

type ACL struct {
	AppliedEntity EntityInterface
	TrafficDerank AutoRateLimit // sets the post to de-rank from search if it gets too hot
	LimitedEdits  RateLimit     // the rate limit of the post
	AllowDiscover bool          // allow the post to be discovered
	AllowView     bool
	AllowEdit     bool
	AllowReact    bool
}
