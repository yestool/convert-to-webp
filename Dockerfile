# 使用官方的Golang镜像作为基础镜像
FROM golang:1.21-alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories 
# 安装libvips依赖
RUN apk add --no-cache vips-dev pkgconfig gcc musl-dev

# 设置工作目录
WORKDIR /app

ENV GO111MODULE=on \
  CGO_ENABLED=1 \
  GOOS=linux \
  GOPROXY="https://goproxy.io,direct"

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目代码
COPY . .

# 构建项目
RUN go build -o image-converter .

# 使用轻量级的Alpine镜像作为运行时镜像
FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories 
# 安装libvips运行时依赖
RUN apk add --no-cache vips

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/image-converter .

# 暴露端口
EXPOSE 8080

# 运行服务
CMD ["./image-converter"]