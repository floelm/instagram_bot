package routines

import (
	"context"
	"errors"
	"github.com/chromedp/chromedp"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/actions"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/cache"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/setup"
	"sync"
	"time"
)

type UnfollowRoutine struct {
	userCache cache.UserCache
	userUrl   chan string
}

func NewUnfollowRoutine() UnfollowRoutine {
	userChannel := make(chan string, 0)

	return UnfollowRoutine{userCache: cache.NewUserCache(), userUrl: userChannel}
}

func (r *UnfollowRoutine) Run() {
	var wg sync.WaitGroup
	wg.Add(2)

	go r.FindUsersToUnfollow()
	go r.UnfollowUser()

	wg.Wait()
}

func (r *UnfollowRoutine) FindUsersToUnfollow() {
	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			println("fetching users")
			now := time.Now()
			keys := r.userCache.GetAllKeys()

			for _, key := range keys {
				expiresAtStamp := r.userCache.Get(key)
				expiresAt := time.Unix(expiresAtStamp, 0)

				if expiresAt.Before(now) {
					r.userUrl <- key
				}
			}
		}
	}
}

func (r *UnfollowRoutine) UnfollowUser() {
	for {
		msg, ok := <-r.userUrl

		if !ok {
			panic(errors.New("failed to unfollow user"))
		}

		println("deleting user from cache")
		r.userCache.Delete(msg)

		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			setup.NonHeadless,
		)

		allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()

		ctx, cancel := chromedp.NewContext(allocCtx)
		defer cancel()

		// create a timeout
		ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
		defer cancel()

		err := actions.RunWrap(ctx,
			actions.PerformLogin(),
			actions.GetDelay(),
			actions.UnfollowUser(msg),
		)

		if err != nil {
			panic(err)
		}
	}
}
