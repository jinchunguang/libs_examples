// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "1050,950"),
		chromedp.UserDataDir(dir),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()


	// create a timeout
	ctx, cancel = context.WithTimeout(taskCtx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err = chromedp.Run(ctx,
		chromedp.Navigate(`https://etax.shandong.chinatax.gov.cn/enterprise/dzswjlogin/dzswj_login.jsp?type=wybs`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > OBJECT`),
		// find and click "Expand All" link
		chromedp.Click(`#zrrLogin`, chromedp.NodeVisible),
		// retrieve the value of the textarea
		chromedp.Value(`#userId`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}