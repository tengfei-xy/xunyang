# xunyang.site - 我的静态生成的博客
从网易博客到wordpress.com,后来wordpress.org，到今天。这一路记录下的点点滴滴算是见证了我的技术发展。
如今我已然具备写网站写前后的能力，受到github.io和jekyll的启发决定写一个**静态生成博客**。

# 额外文件
1. web/docs存放你的.md文件(博客会忽略.开头的隐藏文件)
2. ssl/下存放证书文件(.crt和.pub)

# 启动方式

## 源码启动
```
go env -w GOPROXY=https://goproxy.cn
go env -w GO111MODULE=on
go run main.go
```

## 二进制启动
linux 64位下:
`sudo ./main`

# 其他说明
（在main.go文件的xyWebRun函数中）默认开启80和ssl的443端口

# 开发记录
2020.7.30 重置Centos,安装go14.6,git init
2020.8.1 完成主要功能，查看博文
2020.10.06 小改动，完善README.MD
