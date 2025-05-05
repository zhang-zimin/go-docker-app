# 使用官方的Go基础镜像
FROM golang:1.21

# 设置工作目录
WORKDIR /app

# 复制当前目录下的所有文件到工作目录
COPY . .

# 构建Go应用
RUN go build -o main .

# 暴露8080端口
EXPOSE 8080

# 定义容器启动时执行的命令
CMD ["./main"]
