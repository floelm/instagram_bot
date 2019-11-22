package actions

import (
	"context"
	"github.com/chromedp/chromedp"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/setup"
)

const (
	email    = "test187@web.de" // "blzdontblockus@web.de"
	password = "Hallo123456!"
)

func PerformLogin(ctx context.Context) error {
	err := setup.RunWrap(ctx,
		chromedp.Navigate(`https://www.instagram.com/accounts/login/?source=auth_switcher`),
		chromedp.SendKeys(`input[name="username"]`, email, chromedp.NodeVisible),
		GetDelay(),
		chromedp.SendKeys(`input[name="password"]`, password, chromedp.NodeVisible),
		GetDelay(),
		chromedp.RemoveAttribute(`button[type="submit"]`, "disabled"),
		chromedp.Click(`button[type="submit"]`, chromedp.NodeVisible),
		GetDelay(),
		chromedp.Click(`.aOOlW.HoLwm`, chromedp.NodeVisible),
		GetDelay(),
	)

	return err
}
