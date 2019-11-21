package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var Headers = map[string]interface{}{
	"accept-language": "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"accept-encoding": "gzip, deflate, br",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
	"cache-control":   "no-cache",
	"user-agent":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.50 Safari/537.36",
}

const JsSetupScript = `(function(w, n, wn) {
  // Pass the Webdriver Test.
  Object.defineProperty(n, 'webdriver', {
    get: () => false,
  });

  // Pass the Plugins Length Test.
  // Overwrite the plugins property to use a custom getter.
  Object.defineProperty(n, 'plugins', {
    // This just needs to have length > 0 for the current test,
    // but we could mock the plugins too if necessary.
    get: () => [1, 2, 3, 4, 5],
  });

  // Pass the Languages Test.
  // Overwrite the plugins property to use a custom getter.
  Object.defineProperty(n, 'languages', {
    get: () => ['en-US', 'en'],
  });

  // Pass the Chrome Test.
  // We can mock this in as much depth as we need for the test.
  w.chrome = {
    runtime: {},
  };

  // Pass the Permissions Test.
  const originalQuery = wn.permissions.query;
  return wn.permissions.query = (parameters) => (
    parameters.name === 'notifications' ?
      Promise.resolve({ state: Notification.permission }) :
      originalQuery(parameters)
  );

})(window, navigator, window.navigator);`

const (
	hashtag = "#vegan"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		NonHeadless,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 1000*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		PerformLogin()...
	)

	if err != nil {
		log.Fatal(err)
	}

	err = chromedp.Run(ctx,
		chromedp.SendKeys(`input[placeholder="Suchen"]`, hashtag, chromedp.NodeVisible),
		GetDelay(),
		chromedp.SendKeys(`input[placeholder="Suchen"]`, kb.Enter+kb.Enter, chromedp.NodeVisible),
		GetDelay(),
		chromedp.Click("article > div > div > div > div > div > a > div", chromedp.NodeVisible),
		GetDelay(),
		chromedp.Click(`article section button[type="button"]`, chromedp.NodeVisible),
		GetDelay(),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = chromedp.Run(ctx,
		FollowFirstXInList(20)...,
		//FollowNumberInList(3),
		//chromedp.Click(`/html/body/div[4]/div/div[2]/div/div/div[3]/div[3]/button`, chromedp.NodeVisible, chromedp.BySearch),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func NonHeadless(a *chromedp.ExecAllocator) {
	chromedp.Flag("headless", false)(a)
	chromedp.Flag("hide-scrollbars", false)(a)
	chromedp.Flag("mute-audio", true)(a)
}

func GetDelay() chromedp.Action {
	min := 1
	max := 2
	randomInt := rand.Intn(max-min) + min
	return chromedp.Sleep(time.Duration(randomInt) * time.Second)
}

func PerformLogin() []chromedp.Action {
	return []chromedp.Action{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(Headers)),

		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			_, err = page.AddScriptToEvaluateOnNewDocument(JsSetupScript).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),

		chromedp.Navigate(`https://www.instagram.com/accounts/login/?source=auth_switcher`),
		chromedp.SendKeys(`input[name="username"]`, `blzdontblockus@web.de`, chromedp.NodeVisible),
		GetDelay(),
		chromedp.SendKeys(`input[name="password"]`, `Hallo123456!`, chromedp.NodeVisible),
		GetDelay(),
		chromedp.RemoveAttribute(`button[type="submit"]`, "disabled"),
		chromedp.Click(`button[type="submit"]`, chromedp.NodeVisible),
		GetDelay(),
		chromedp.Click(`.aOOlW.HoLwm`, chromedp.NodeVisible),
		GetDelay(),
	}
}

func FindItemFromSearch(hashtag string) []chromedp.Action {
	return []chromedp.Action{
		chromedp.SendKeys(`input[placeholder="Suchen"]`, hashtag, chromedp.NodeVisible),
		GetDelay(),
		chromedp.SendKeys(`input[placeholder="Suchen"]`, kb.Enter+kb.Enter, chromedp.NodeVisible),
		GetDelay(),
		chromedp.Click("article > div > div > div > div > div > a > div", chromedp.NodeVisible),
		GetDelay(),
		chromedp.Click(`article section button[type="button"]`, chromedp.NodeVisible),
		GetDelay(),
	}
}

func FollowFirstXInList(count int) []chromedp.Action {
	actions := make([]chromedp.Action, 0)

	for i := 1; i < count; i++ {
		actions = append(actions, GetDelay())
		actions = append(actions, FollowNumberInList(i))
	}

	return actions
}

func FollowNumberInList(number int) chromedp.Action {
	return chromedp.Click(`/html/body/div[4]/div/div[2]/div/div/div[`+ strconv.Itoa(number) +`]/div[3]/button`, chromedp.NodeVisible, chromedp.BySearch)
}
