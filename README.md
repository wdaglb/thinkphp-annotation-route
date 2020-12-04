## 使用

下载kroute文件放到任意目录,设置环境变量到该目录即可

命令行运行kroute解析注释路由。

> 监听文件改动 --watch

```
# 监听文件改动并解析路由
kroute --watch

# 解析路由
kroute

# 如果你的app目录不是application，可使用-app=apps来指定
kroute -app=apps
```

## 源代码编译

```
CGO_ENABLED=0;GOOS=linux
or
CGO_ENABLED=0;GOOS=windows

go build main.go -o bin/kroute
```