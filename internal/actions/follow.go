package actions

import (
	"context"
	"github.com/chromedp/chromedp"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/setup"
	"strconv"
)

func UnfollowUser(ctx context.Context, name string) error {
	err := setup.RunWrap(ctx,
		chromedp.Navigate(`https://www.instagram.com/`+name),
		GetDelay(),
		GetDelay(),
		GetDelay(),
		chromedp.Click(`/html/body/span/section/main/div/header/section/div[1]/div[1]/span/span[1]/button`, chromedp.NodeVisible),
		GetDelay(),
		GetDelay(),
		GetDelay(),
		chromedp.Click(`/html/body/div[3]/div/div/div[3]/button[1]`, chromedp.NodeVisible),
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
		GetDelay(),
		chromedp.Click(`/html/body/div[4]/div/div[2]/div/div/div[`+strconv.Itoa(number)+`]/div[3]/button`, chromedp.NodeVisible, chromedp.BySearch),
	)

	return err
}
