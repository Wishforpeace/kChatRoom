###KChatRoom在线多人聊天室

使用Websocket和Gin框架基于Golang开发的在线聊天室

Github项目地址： [地址](https://github.com/linkaias/kChatRoom)

详细说明介绍文档：[地址](https://www.uiucode.com/view/41.html)

项目在线体验地址：[地址](http://kchatroom.uiucode.com)
>（可使用自己的邮箱注册也可使用体验邮箱：用户名：user@qq.com 密码：123456  体验邮箱2：用户名：user2@qq.com 密码：123456）系统只允许单点登录，体验邮箱登陆后可能被别人挤掉。

#### 项目功能：
1. 简单聊天机器人功能，用户上线欢迎。 
2. 自定义头像捏脸功能。 
3. 类似QQ消息声音提醒功能。
4. 保存聊天记录功能，上拉聊天记录会分页拉取最近的聊天记录。
5. 实现用户单点登录，统一账户仅支持单点登录。
6. 登陆页面在线人数展示。
7. 表情系统，可发送多个丰富表情。
8. 基于Redis实现tcp用户安全登陆。

#### 配置文件更改：
1. 复制项目根目录下config/config_bak.yml 为config.yml。
2. 按需修改config.yml中配置。
3. 如果搭建域名使用，需要在配置文件中把cookie下域名配置改为你的域名。
####项目基础运行环境：
1. Golang1.12+
2. Redis
3. Mysql
####项目安装使用说明：
1. 从Github上下载项目https://github.com/linkaias/kChatRoom
2. 导入数据库文件，用于储存聊天记录和用户信息。
3. 运行项目：在项目根目录 go run main.go
4. 浏览器打开：http://127.0.0.1:8060/
5. 登陆：用户名：user 密码：123456  或者：用户名：user2 密码：123456

说明：此项目为个人学习创建，不足之处还望理解。