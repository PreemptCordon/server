type GroupObj struct {
	handle    string
	parent    GroupObj
	about     Wiki
	region    Region
	languages string
	discovery DiscoverPreferences
}