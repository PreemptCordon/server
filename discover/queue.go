package discover

func DiscoverQueue() {
	Jobs = []DiscoverJobs
	for job in Jobs {
		done := job.evalone()
		if done {
			Jobs.pop(job)
		}
		if Jobs.empty {
			break
		}
	}
}