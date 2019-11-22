package actions

import "github.com/chromedp/chromedp"

func ToggleLike() []chromedp.Action {
	return []chromedp.Action{
		chromedp.Click(`/html/body/div[3]/div[2]/div/article/div[2]/section[1]/span[1]/button`, chromedp.NodeVisible),
		GetDelay(),
	}
}

func Like(attributes []map[string]string) []chromedp.Action {
	return []chromedp.Action{
		chromedp.Click(`/html/body/div[3]/div[2]/div/article/div[2]/section[1]/span[1]/button`, chromedp.NodeVisible),
		GetDelay(),
	}
}

func UnLike(attributes []map[string]string) []chromedp.Action {
	return []chromedp.Action{
		chromedp.Click(`/html/body/div[3]/div[2]/div/article/div[2]/section[1]/span[1]/button`, chromedp.NodeVisible),
		GetDelay(),
	}
}