// Command emulate is a chromedp example demonstrating how to emulate a
// specific device.
package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

func main() {
	// create context
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	// 启动可视化
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false),
		chromedp.Flag("ignore-certificate-errors", true),
		//chromedp.Flag("window-size", "1050,950"),
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	// run
	var b1, b2 []byte
	if err := chromedp.Run(ctx,
		// 模拟iPhone7环境
		chromedp.Emulate(device.IPhone7),
		chromedp.Navigate(`https://www.baidu.com/`),
		chromedp.CaptureScreenshot(&b1),

		// reset
		chromedp.Emulate(device.Reset),

		// 设置大的浏览窗口
		chromedp.EmulateViewport(1920, 2000),
		chromedp.Navigate(`https://www.baidu.com/`),
		chromedp.CaptureScreenshot(&b2),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("screenshot_iPhone7.png", b1, 0777); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("screenshot_pc.png", b2, 0777); err != nil {
		log.Fatal(err)
	}
}
