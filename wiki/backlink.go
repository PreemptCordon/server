package wiki

import "fmt"

func request_backlink(source_requester string, cites_approver string) {
	fmt.Println(source_requester, cites_approver)
	//cites_approver.backlinks.add(source_requester, states.pending)
}
func approve_backlink(cites_approver string, source_requester string) {
	fmt.Println(source_requester, cites_approver)
	//cites_approver.backlinks.add(source_requester, states.approved)
}
