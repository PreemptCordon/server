package wiki

func request_backlink(source_requester, cites_approver) {
	cites_approver.backlinks.add(source_requester, states.pending)
}
func approve_backlink(cites_approver, source_requester) {
	cites_approver.backlinks.add(source_requester, states.approved)
}
