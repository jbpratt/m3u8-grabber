package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventResponseReceived:
			fmt.Printf("Response: %s\n", ev.Response.URL)
		}
	})

	chromedp.Run(ctx, network.Enable())
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.livenewswatch.com/cnn-news-usa.html"),
		chromedp.WaitVisible(".jw-text-live", chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	}

}
