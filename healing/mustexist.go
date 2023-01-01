func mustexist() {
	/* job to always ensure:
	- a server group
	- a server TOS wikki
	- a server-admin
	*/
	if !db.groups["site"] {
		db.groups["site"] = demo_group
	}
	if !db.wiki["TOS"] {
		db.wiki["TOS"] = demo_tos
	}
	if db.roles.siteadmin == nil {
		db.roles.siteadmin.append(db.users.highestkarma)
	}
}