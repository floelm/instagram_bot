package main

import (
	"gitlab.applike-services.info/mcoins/backend/insta/internal/routines"
	"sync"
)

func main() {
	unfollowRoutine := routines.NewUnfollowRoutine()
	followRoutine := routines.NewFollowRoutine()
	var wg sync.WaitGroup
	wg.Add(2)

	go unfollowRoutine.Run()
	go followRoutine.Run()

	wg.Wait()
}
