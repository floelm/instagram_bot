package routines

import (
	"context"
	"github.com/chromedp/chromedp"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/actions"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/setup"
	"log"
	"sync"
	"time"
)

type CommentRoutine struct {
	userUrl chan string
}

func NewCommentRoutine() CommentRoutine {
	userChannel := make(chan string, 0)

	return CommentRoutine{userUrl: userChannel}
}

func (r *CommentRoutine) Run() {
	var wg sync.WaitGroup
	wg.Add(2)

	go r.comment(10*time.Second, "#cat")

	wg.Wait()
}

func (r *CommentRoutine) comment(interval time.Duration, hashtag string) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				setup.NonHeadless,
			)

			allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
			defer cancel()

			ctx, cancel := chromedp.NewContext(allocCtx)
			defer cancel()

			// create a timeout
			ctx, cancel = context.WithTimeout(ctx, 1000*time.Second)
			defer cancel()

			// abuse this var
			var err error

			setup.SetupClient(ctx)
			err = actions.PerformLogin(ctx)

			err = FindItemFromSearch(ctx, hashtag)

			err = OpenPostOnDiscovery(ctx, 1)
			err = actions.Comment(ctx)

			if err != nil {
				log.Fatal(err)
			}

			cancel()
		}
	}
}
