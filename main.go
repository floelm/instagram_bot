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

	err = OpenPostOnDiscorvery(ctx, 1)

	err = actions.UnLike(ctx)

	err = setup.RunWrap(ctx,
		actions.GetDelay(),
		chromedp.Click(`article section button[type="button"]`, chromedp.NodeVisible),
		actions.GetDelay(),
	)

	err = FollowFirstXInList(ctx, 1)

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

func OpenPostOnDiscorvery(ctx context.Context, position int) error {
	err := setup.RunWrap(ctx,
		actions.GetDelay(),
		chromedp.Click(`//*[@id="react-root"]/section/main/article/div[1]/div/div/div[`+strconv.Itoa(position)+`]/div[2]/a/div`, chromedp.NodeVisible),
		actions.GetDelay(),
	)

	return err
}

func FollowFirstXInList(ctx context.Context, count int) error {
	var err error

	for i := 2; i < count; i++ {
		err = FollowNumberInList(ctx, i)
	}

	return err
}

func FollowNumberInList(ctx context.Context, number int) error {
	err := setup.RunWrap(ctx,
		actions.GetDelay(),
		chromedp.Click(`/html/body/div[4]/div/div[2]/div/div/div[`+strconv.Itoa(number)+`]/div[3]/button`, chromedp.NodeVisible, chromedp.BySearch),
	)

	return err
}
