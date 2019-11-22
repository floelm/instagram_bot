package actions

import (
	"github.com/chromedp/chromedp"
	"math/rand"
	"time"
)

func GetDelay() chromedp.Action {
	min := 1
	max := 2
	randomInt := rand.Intn(max-min) + min
	return chromedp.Sleep(time.Duration(randomInt) * time.Second)
}

func GetClassAttribute(attributeSlice []map[string]string) *string {
	for _, a := range attributeSlice {
		classValue := a["class"]
		return &classValue
	}

	return nil
}
