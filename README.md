本代码为李文周老师web开发进阶项目实战实现的bulebell后端源码

环境变量： Environment正确填写内容：GOPROXY=https://goproxy.cn,direct
加载模块：go mod tidy
折叠函数： command + shift + -
快捷复制行：command + d
jwt是基于token的认证模式，与cookies sessions模式最大的区别就是不用再服务端存储认证数据
air 热加载
go install github.com/cosmtrek/air@latest

go env GOPATH   查看假设输出的路径为  /a/b

vim ~/.zshrc   添加  alias air='/a/b/air'

我的是 alias air='/Users/kaapo/go/bin/air'

source  ~/.zshrc

air  -v    查看是否成功