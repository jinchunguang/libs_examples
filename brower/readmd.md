``` 
chromedp.BySearch // 如果不写，默认会使用这个选择器，类似devtools ctrl+f 搜索
chromedp.ByID // 只id来选择元素
chromedp.ByQuery // 根据document.querySelector的规则选择元素，返回单个节点
chromedp.ByQueryAll // 根据document.querySelectorAll返回所有匹配的节点
chromedp.ByNodeIP // 检索特定节点(必须先有分配的节点IP)，这个暂时没用过也没看到过例子，如果有例子可以发给我看下


chromedp.Navigate("https://xxxx") // 设置url
chromedp.WaitVisible(`#username`, chromedp.ByID), //  使用chromedp.ByID选择器。所以就是等待id=username的标签元素加载完。
chromedp.SendKeys(`#username`, "username", chromedp.ByID), // 使用chromedp.ByID选择器。向id=username的标签输入username。
chromedp.Value(`#input1`, val1, chromedp.ByID), // 获取id=input1的值，并把值传给val1
chromedp.Click("btn-submit",chromedp.Bysearch), // 触发点击事件，
chromedp.Screenshot(`#row`, &buf, chromedp.ByID), // 截图id=row的标签，把值传入buf 需要事先定义var buf []byte
chromedp.ActionFunc(func(context.Context, cdp.Executor) error { // 将图片写入文件
return ioutil.WriteFile("test.png", buf, 0644)
```