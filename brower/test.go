package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

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

	taskCtx, cancel = context.WithTimeout(taskCtx, 10*time.Second)
	defer cancel()

	// 确保浏览器进程已启动
	if err := chromedp.Run(taskCtx); err != nil {
		panic(err)
	}

	// listen network event
	listenForNetworkEvent(taskCtx)

	chromedp.Run(taskCtx,
		network.Enable(),
		chromedp.Navigate(`https://etax.shandong.chinatax.gov.cn/enterprise/dzswjlogin/dzswj_login.jsp?type=main`),
		chromedp.WaitVisible(`body`, chromedp.BySearch),
	)

}

//监听
func listenForNetworkEvent(ctx context.Context) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventResponseReceived:
			resp := ev.Response
			if len(resp.Headers) != 0 {
				//log.Printf("received headers: %s", resp.Headers)
				log.Printf("received Status: %s", resp.Status)

				if strings.Index(resp.URL, ".ts") != -1 {
					log.Printf("received headers: %s", resp.URL)
				}
			}

		}
		// other needed network Event
	})
}