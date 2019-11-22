package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/actions"
	"log"
	"strconv"
	"time"
)

const (
	hashtag = "#vegan"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		actions.NonHeadless,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 1000*time.Second)
	defer cancel()

	attributes := make([]map[string]string, 0)

	err := actions.RunWrap(ctx,
		actions.PerformLogin(),
		FindItemFromSearch(hashtag),
		OpenPostOnDiscorvery(1),
		actions.Like(attributes),
	)

	err = actions.RunWrap(ctx,
		chromedp.Sleep(1000*time.Second),
		actions.GetDelay(),
		chromedp.Click(`article section button[type="button"]`, chromedp.NodeVisible),
		actions.GetDelay(),
		FollowFirstXInList(1),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func FindItemFromSearch(hashtag string) []chromedp.Action {
	return []chromedp.Action{
		chromedp.SendKeys(`input[placeholder="Suchen"]`, hashtag, chromedp.NodeVisible),
		actions.GetDelay(),
		chromedp.SendKeys(`input[placeholder="Suchen"]`, kb.Enter+kb.Enter, chromedp.NodeVisible),
		actions.GetDelay(),
	}
}

func OpenPostOnDiscorvery(position int) []chromedp.Action {
	return []chromedp.Action{
		actions.GetDelay(),
		chromedp.Click(`//*[@id="react-root"]/section/main/article/div[1]/div/div/div[`+strconv.Itoa(position)+`]/div[2]/a/div`, chromedp.NodeVisible),
		actions.GetDelay(),
	}
}

func FollowFirstXInList(count int) []chromedp.Action {
	actionsToExecute := make([]chromedp.Action, 0)

	for i := 2; i < count; i++ {
		actionsToExecute = append(actionsToExecute, actions.GetDelay())
		actionsToExecute = append(actionsToExecute, FollowNumberInList(i))
	}

	return actionsToExecute
}

func FollowNumberInList(number int) chromedp.Action {
	return chromedp.Click(`/html/body/div[4]/div/div[2]/div/div/div[`+strconv.Itoa(number)+`]/div[3]/button`, chromedp.NodeVisible, chromedp.BySearch)
}
