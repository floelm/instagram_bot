package main

import (
	"gitlab.applike-services.info/mcoins/backend/insta/internal/routines"
	"sync"
)

func main() {
	unfollowRoutine := routines.NewUnfollowRoutine()
	followRoutine := routines.NewFollowRoutine()
	commentRoutine := routines.NewCommentRoutine()
	var wg sync.WaitGroup
	wg.Add(3)

	go unfollowRoutine.Run()
	go followRoutine.Run()
	go commentRoutine.Run()

	wg.Wait()
}
