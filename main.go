package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
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

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		NonHeadless,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
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
		chromedp.Sleep(2 * time.Second),
		chromedp.SendKeys(`input[name="password"]`, `Hallo123456!`, chromedp.NodeVisible),
		chromedp.Sleep(2 * time.Second),
		chromedp.RemoveAttribute(`button[type="submit"]`, "disabled"),
		chromedp.Click(`button[type="submit"]`, chromedp.NodeVisible),
		chromedp.Sleep(5 * time.Second),
		chromedp.Click(`button`, chromedp.NodeVisible),
	)

/*	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://golang.org/pkg/time/`),
		chromedp.Text(`#pkg-overview`, &example, chromedp.NodeVisible, chromedp.ByID),
	)*/

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}

func NonHeadless(a *chromedp.ExecAllocator) {
	chromedp.Flag("headless", false)(a)
	chromedp.Flag("hide-scrollbars", false)(a)
	chromedp.Flag("mute-audio", true)(a)
}
