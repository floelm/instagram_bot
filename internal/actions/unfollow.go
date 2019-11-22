package actions

import "github.com/chromedp/chromedp"

func UnfollowUser(profileUrl string) []chromedp.Action {
	return []chromedp.Action{
		chromedp.Navigate(profileUrl),
		GetDelay(),
		chromedp.Click(`/html/body/span/section/main/div/header/section/div[1]/div[1]/span/span[1]/button`, chromedp.NodeVisible),
		GetDelay(),
		chromedp.Click(`/html/body/div[3]/div/div/div[3]/button[1]`, chromedp.NodeVisible),
	}
}
