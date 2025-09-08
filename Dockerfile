# 构建阶段
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o ipaas-agent .

# 运行阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/ipaas-agent .
# 创建config目录并复制配置文件
RUN mkdir -p config
COPY config/config.yml config/

# 暴露端口（根据实际需要调整）
EXPOSE 8080

# 设置容器启动命令
CMD ["./ipaas-agent"]