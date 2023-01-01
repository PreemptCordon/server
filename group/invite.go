package group

func inviteExtend(host UserObj, guest UserObj, group GroupObj){
	group.invites.add(host=host,guest=guest)
	notify(guest,"you have been invited to {group} by {host}")
}
func inviteAccept(guest UserObj, invite) bool {
	valid, host = group.invites.validate(invite)
	if ! valid {
		return false
	}
	group.members.add(guest)
	host.karma.link(guest)
}
func inviteRevoke(user UserObj, group GroupObj, invite){
	group.invites.remove(invite)
}