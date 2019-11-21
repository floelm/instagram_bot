package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"math/rand"
	"time"
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
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.applike.info/`),
		//chromedp.WaitVisible(`body`),
		chromedp.Text(`#layers-widget-column-51-443`, &example, chromedp.NodeVisible, chromedp.ByID),
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

func GetDelay() {
	min := 1
	max := 5
	randomInt := rand.Intn(max-min) + min
	chromedp.Sleep(time.Duration(randomInt) * time.Second)
}
