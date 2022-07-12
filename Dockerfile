# FORM 表示你的镜像所依托的基础镜像是什么，我们选择golang,latest表示最新版本
FROM golang:latest
# 写一下作者信息
MAINTAINER zxp@1503450771@qq.com

# 使用代理,环境变量， 如果不设置goproxy，你在使用go get命令拉取go库的时候大概率会出现超时情况，拉取失败！
ENV GOPROXY https://goproxy.cn,direct

# 在 GOPATH 指定的工作目录下，代码总是会保存在 $GOPATH/src 目录下。
# 在工程经过 go build、go install 或 go get 等指令后，会将产生的二进制可执行文件放在 $GOPATH/bin 目录下，生成的中间缓存文件会被保存在 $GOPATH/pkg 下。
WORKDIR $GOPATH/usr/local/project/myblog

# 把当前目录拷贝到对应的镜像目录
COPY . $GOPATH/usr/local/project/myblog

# 制作镜像操作指令
RUN go build .

# 使用端口号
EXPOSE 8889

# 容器启动过后执行的命令，一般我们使用的go是最新版本，需要go mod tidy 与项目中使用的版本相对应
ENTRYPOINT ["go","mod","tidy"]

# 启动程序
ENTRYPOINT ["./myblog"]