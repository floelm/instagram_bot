package actions

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"math/rand"
	"reflect"
	"time"
)

func RunWrap(ctx context.Context, actions ...interface{}) error {
	actionsToExecute := make([]chromedp.Action, 0)

	for _, a := range actions {
		rt := reflect.TypeOf(a)

		switch rt.Kind() {
		case reflect.Slice:
			actionsToExecute = splitActionsInto(actionsToExecute, a.([]chromedp.Action))
		default:
			actionsToExecute = append(actionsToExecute, a.(chromedp.Action))
		}
	}

	err := chromedp.Run(ctx,
		actionsToExecute...,
	)

	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}

func GetDelay() chromedp.Action {
	min := 1
	max := 2
	randomInt := rand.Intn(max-min) + min
	return chromedp.Sleep(time.Duration(randomInt) * time.Second)
}

func splitActionsInto(actionsToExecute []chromedp.Action, actions []chromedp.Action) []chromedp.Action {
	return append(actionsToExecute, actions...)
}
