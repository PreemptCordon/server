package main

var visits int
loadwatch := make(chan int)
func loadmiddleware(){
	atomic.Add(&visits, 1)
}

func emitresetload(){
	// once a second, emit the result of load and reset to zero
	loadwatch.push(visits)
	visits = 0
}

go func loadlisten(){
	load := <- loadwatch
	if load < MAX_BACKGROUND {
		DiscoverQueue.runonce()
	}
}
