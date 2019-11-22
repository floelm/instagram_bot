package actions

import (
	"context"
	"github.com/chromedp/chromedp"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/setup"
)

func UnfollowUser(ctx context.Context, profileUrl string) error {
	err := setup.RunWrap(ctx,
		chromedp.Navigate(profileUrl),
		GetDelay(),
		chromedp.Click(`/html/body/span/section/main/div/header/section/div[1]/div[1]/span/span[1]/button`, chromedp.NodeVisible),
		GetDelay(),
		chromedp.Click(`/html/body/div[3]/div/div/div[3]/button[1]`, chromedp.NodeVisible),
	)

	return err
}
