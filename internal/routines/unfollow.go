package routines

import (
	"gitlab.applike-services.info/mcoins/backend/insta/internal/cache"
	"time"
)

type UnfollowRoutine struct {
	userCache cache.UserCache
	userUrl   chan string
}

func NewUnfollowRoutine() UnfollowRoutine {
	return UnfollowRoutine{userCache: cache.NewUserCache()}
}

func (r *UnfollowRoutine) Run() {

}

func (r *UnfollowRoutine) FindUsersToUnfollow() {
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

func (r *UnfollowRoutine) UnfollowUser() {
	select {
	case userUrl := <-r.userUrl:
		r.userCache.Delete(userUrl)
		//start chrome
	}
}
