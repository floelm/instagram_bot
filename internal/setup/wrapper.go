package setup

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"reflect"
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

func splitActionsInto(actionsToExecute []chromedp.Action, actions []chromedp.Action) []chromedp.Action {
	return append(actionsToExecute, actions...)
}
