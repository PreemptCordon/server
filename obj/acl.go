package obj

type ACL struct {
	Entity  EntityObj
	Traffic AutoRateLimit
	Limited RateLimit
	View    bool
	Edit    bool
	React   bool
}
