package main

import (
	"gitlab.applike-services.info/mcoins/backend/insta/internal/routines"
	"sync"
)

func main() {
	routine := routines.NewUnfollowRoutine()
	var wg sync.WaitGroup
	wg.Add(1)

	go routine.Run()

	wg.Wait()
}
