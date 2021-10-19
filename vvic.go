package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)
func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoDefaultBrowserCheck,                   //不检查默认浏览器
		chromedp.Flag("headless", false),                 //开启图像界面
		chromedp.Flag("ignore-certificate-errors", true), //忽略错误
		chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
		chromedp.Flag("disable-extensions", false),  //开启插件支持
		chromedp.Flag("disable-default-apps", false),
		chromedp.NoFirstRun, //不是首次运行
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	//allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	//defer cancel()
	//标签1
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.vvic.com/item/61262ddc09e31c00016b5bde"),
		//chromedp.WaitVisible(`body > div.deatil-wrapper > article > main > section:nth-child(1) > div.detail.product-detail > div.detail-price > dl:nth-child(1) > dd`, chromedp.ByQuery),
		chromedp.Text(".detail-price", &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}