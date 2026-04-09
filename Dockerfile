# 第一阶段：编译阶段 (使用大的基础镜像)
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# 第二阶段：运行阶段 (使用极小的基础镜像，只放编译好的二进制文件)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]