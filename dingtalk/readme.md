
> 钉钉群机器人定义

群机器人是钉钉群的高级扩展功能。群机器人可以将第三方服务的信息聚合到群聊中，实现自动化的信息同步。目前，大部分机器人在添加后，还需要进行Webhook配置，才可正常使用(配置说明详见操作流程中的帮助链接)。

例如：

通过聚合GitHub，GitLab等源码管理服务，实现源码更新同步。

通过聚合Trello，JIRA等项目协调服务，实现项目信息同步。

另外，群机器人支持Webhook协议的自定义接入，支持更多可能性，例如：你可将运维报警通过自定义机器人聚合到钉钉群实现提醒功能。

>  机器人发送消息频率限制

消息发送太频繁会严重影响群成员的使用体验，大量发消息的场景（譬如系统监控报警）可以将这些信息进行整合，通过markdown消息以摘要的形式发送到群里。

`每个机器人每分钟最多发送20条。如果超过20条，会限流10分钟。`

>  PC客户端配置入口
登录钉钉PC客户端(请升级至最新版)，操作入口：

窗口右上角点击【头像】，进入”机器人管理“， 可对所有机器人进行统一管理。
进入一个钉钉群，在群的顶部功能栏中，点击【群设置】，进入菜单可以看到【群机器人】的入口，点击进入“群机器人”的管理面板后，可以进行添加、编辑和删除群机器人的操作。

![image.png](https://upload-images.jianshu.io/upload_images/1779921-8d397b4dea70a3ba.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

**我们这里使用最后一个webhook**

>   获取自定义机器人webhook

步骤一，在机器人管理页面选择“自定义”机器人，输入机器人名字并选择要发送消息的群。如果需要的话，可以为机器人设置一个头像。点击“完成添加”，完成后会生成Hook地址，如下图：

![image](https://upload-images.jianshu.io/upload_images/1779921-3494a7128fabfb08?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

步骤二，点击“复制”按钮，即可获得这个机器人对应的Webhook地址，其格式如下：

```
https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxx
```
>  使用自定义机器人

（1）获取到Webhook地址后，用户可以向这个地址发起HTTP POST 请求，即可实现给该钉钉群发送消息。注意，发起POST请求时，必须将字符集编码设置成UTF-8。

（2）当前自定义机器人支持文本 (text)、链接 (link)、markdown(markdown)、ActionCard、FeedCard消息类型，大家可以根据自己的使用场景选择合适的消息类型，达到最好的展示样式。

（3）自定义机器人发送消息时，可以通过手机号码指定“被@人列表”。在“被@人列表”里面的人员收到该消息时，会有@消息提醒(免打扰会话仍然通知提醒，首屏出现“有人@你”)。

（4）当前机器人尚不支持应答机制 (该机制指的是群里成员在聊天@机器人的时候，钉钉回调指定的服务地址，即Outgoing机器人)。


>  php 代码

```php

<?php
/**
 * php 使用钉钉机器人发送消息.
 * User: Administrator
 * Date: 2019/8/29
 * Time: 23:40
 */

function request_by_curl($remote_server, $post_string)
{
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $remote_server);
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 5);
    curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type: application/json;charset=utf-8'));
    curl_setopt($ch, CURLOPT_POSTFIELDS, $post_string);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    // 不用开启curl证书验证
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
    $data = curl_exec($ch);
    //$info = curl_getinfo($ch);
    //var_dump($info);
    curl_close($ch);
    return $data;
}

$webhook = "https://oapi.dingtalk.com/robot/send?access_token=xxxxxx";

// text类型
$textString = json_encode([
    'msgtype' => 'text',
    'text' => [
        "content" => "我就是我, 是不一样的烟火@156xxxx8827"
    ],
    'at' => [
        'atMobiles' => [
            "156xxxx8827",
            "189xxxx8325"
        ],
        'isAtAll' => false

    ]
]);

// link类型
$textString = json_encode([
    "msgtype" => "link",
    "link" => [
        "text" => "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。
而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？",
        "title" => "时代的火车向前开",
        "picUrl" => "",
        "messageUrl" => "https://www.dingtalk.com/",
    ]
]);

// markdown类型
$textString = json_encode([
    "msgtype" => "markdown",
    "markdown" => [
        "title" => "杭州天气",
        "text" => "#### 杭州天气 @156xxxx8827\n" .
            "> 9度，西北风1级，空气良89，相对温度73%\n\n" .
            "> ![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)\n" .
            "> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n"
    ],
    "at" => [
        "atMobiles" => [
            "156xxxx8827",
            "189xxxx8325"
        ],
        "isAtAll" => false
    ]
]);

// 整体跳转ActionCard类型
$textString = json_encode([
    "actionCard" => [
        "title" => "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
        "text" => "![screenshot](@lADOpwk3K80C0M0FoA) 
 ### 乔布斯 20 年前想打造的苹果咖啡厅 
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
        "hideAvatar" => "0",
        "btnOrientation" => "0",
        "singleTitle" => "阅读全文",
        "singleURL" => "https://www.dingtalk.com/"
    ],
    "msgtype" => "actionCard"
]);


// 独立跳转ActionCard类型
$textString = json_encode([
    "actionCard" => [
        "title" => "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
        "text" => "![screenshot](@lADOpwk3K80C0M0FoA) 
 ### 乔布斯 20 年前想打造的苹果咖啡厅 
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
        "hideAvatar" => "0",
        "btnOrientation" => "0",
        "btns" => [
            [
                "title" => "内容不错",
                "actionURL" => "https://www.dingtalk.com/"
            ],
            [
                "title" => "不感兴趣",
                "actionURL" => "https://www.dingtalk.com/"
            ]
        ]
    ],
    "msgtype" => "actionCard"
]);

// FeedCard类型
$textString = json_encode([
    "feedCard" => [
        "links" => [
            [
                "title" => "时代的火车向前开1",
                "messageURL" => "https://www.dingtalk.com/",
                "picURL" => "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1567105217584&di=4c91fefc045f54267edcf8c544e3bd3b&imgtype=0&src=http%3A%2F%2Fk.zol-img.com.cn%2Fdcbbs%2F16420%2Fa16419096_s.jpg"
            ],
            [
                "title" => "时代的火车向前开2",
                "messageURL" => "https://www.dingtalk.com/",
                "picURL" => ""
            ]
        ]
    ],
    "msgtype" => "feedCard"
]);


$result = request_by_curl($webhook, $textString);
echo $result;

> 最终效果

```
![image.png](https://upload-images.jianshu.io/upload_images/1779921-90c83f55d13cf646.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/1779921-03c8dc2a3a590cc5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/1779921-827c34806ac27e77.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/1779921-5592cb376a9ee77b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/1779921-0b884cb646666d86.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![image.png](https://upload-images.jianshu.io/upload_images/1779921-89380973ac63e03b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
