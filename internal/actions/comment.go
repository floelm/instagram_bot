package actions

import (
	"context"
	"github.com/chromedp/chromedp"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/setup"
	"math/rand"
)

const (
	CommentTextAreaElement = `/html/body/div[3]/div[2]/div/article/div[2]/section[3]/div[1]/form/textarea`
	CommentFormElement     = `/html/body/div[3]/div[2]/div/article/div[2]/section[3]/div[1]/form/button`
)

var comments = []string{
	"#follow4follow",
	"#like4like",
	"super cool pic :)))))))",
}

func Comment(ctx context.Context) error {
	min := 0
	max := len(comments) - 1
	randomInt := rand.Intn(max-min) + min

	err := setup.RunWrap(ctx,
		chromedp.SendKeys(CommentTextAreaElement, comments[randomInt], chromedp.NodeVisible),
		chromedp.Click(CommentFormElement, chromedp.NodeVisible),
		GetDelay(),
	)

	return err
}
