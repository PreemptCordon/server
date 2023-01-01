package obj

type AutoRateLimit bool

const (
	AutoRateLimitEnabled  AutoRateLimit = true
	AutoRateLimitDisabled AutoRateLimit = false
)

type RateLimit bool

const (
	RateLimitEnabled  RateLimit = true
	RateLimitDisabled RateLimit = false
)
