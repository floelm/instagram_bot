package actions

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

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

var Headers = map[string]interface{}{
	"accept-language": "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
	"accept-encoding": "gzip, deflate, br",
	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
	"cache-control":   "no-cache",
	"user-agent":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.50 Safari/537.36",
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
