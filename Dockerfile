# 使用 golang 官方镜像作为基础镜像
FROM golang:1.16 AS builder
LABEL authors="alan"


# 设置工作目录
WORKDIR /app

# 复制源代码到镜像中
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp

# 使用一个轻量的 alpine 镜像作为最终镜像
FROM alpine:latest
LABEL authors="alan"


# 设置工作目录
WORKDIR /app

# 复制编译后的二进制文件到镜像中
COPY --from=builder /app/myapp .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./myapp"]

ENTRYPOINT ["top", "-b"]