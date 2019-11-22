package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/actions"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/setup"
	"log"
	"strconv"
	"time"
)

const (
	hashtag = "#vegan"
)

func main() {
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

	err = actions.Like(ctx)
	err = actions.Comment(ctx)

	err = setup.RunWrap(ctx,
		chromedp.Sleep(1000*time.Second),
		actions.GetDelay(),
		chromedp.Click(`article section button[type="button"]`, chromedp.NodeVisible),
		actions.GetDelay(),
	)

	err = actions.FollowFirstXInList(ctx, 1)

	err = setup.RunWrap(ctx,
		chromedp.Sleep(100*time.Second),
	)

	if err != nil {
		log.Fatal("upss")
	}
}

func FindItemFromSearch(ctx context.Context, hashtag string) error {
	err := setup.RunWrap(ctx,
		chromedp.SendKeys(`input[placeholder="Suchen"]`, hashtag, chromedp.NodeVisible),
		actions.GetDelay(),
		chromedp.SendKeys(`input[placeholder="Suchen"]`, kb.Enter+kb.Enter, chromedp.NodeVisible),
		actions.GetDelay(),
	)

	return err
}

func OpenPostOnDiscovery(ctx context.Context, position int) error {
	err := setup.RunWrap(ctx,
		actions.GetDelay(),
		chromedp.Click(`//*[@id="react-root"]/section/main/article/div[1]/div/div/div[`+strconv.Itoa(position)+`]/div[2]/a/div`, chromedp.NodeVisible),
		actions.GetDelay(),
	)

	return err
}
