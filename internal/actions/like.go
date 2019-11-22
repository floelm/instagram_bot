package actions

import (
	"context"
	"github.com/chromedp/chromedp"
	"gitlab.applike-services.info/mcoins/backend/insta/internal/setup"
	"strings"
)

const (
	LikeButtonElement     = `/html/body/div[3]/div[2]/div/article/div[2]/section[1]/span[1]/button`
	LikeButtonIconElement = `/html/body/div[3]/div[2]/div/article/div[2]/section[1]/span[1]/button/span`
)

func ToggleLike(ctx context.Context) error {
	err := setup.RunWrap(ctx,
		chromedp.Click(LikeButtonElement, chromedp.NodeVisible),
		GetDelay(),
	)

	return err
}

func Like(ctx context.Context) error {
	isLiked := IsLiked(ctx)

	if isLiked {
		return nil
	}

	err := setup.RunWrap(ctx,
		chromedp.Click(LikeButtonElement, chromedp.NodeVisible),
		GetDelay(),
	)

	return err
}

func UnLike(ctx context.Context) error {
	isLiked := IsLiked(ctx)

	if !isLiked {
		return nil
	}

	err := setup.RunWrap(ctx,
		chromedp.Click(LikeButtonElement, chromedp.NodeVisible),
		GetDelay(),
	)

	return err
}

func IsLiked(ctx context.Context) bool {
	attributes := make([]map[string]string, 0)

	setup.RunWrap(ctx,
		chromedp.AttributesAll(LikeButtonIconElement, &attributes, chromedp.NodeVisible),
		GetDelay(),
	)

	classAttribute := GetClassAttribute(attributes)

	return strings.HasPrefix(*classAttribute, "glyphsSpriteHeart__filled")
}
