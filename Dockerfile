# 构建阶段
FROM golang:1.19-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz

# 生成
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
COPY app.env .
COPY start.sh .
#COPY wait-for.sh .
#RUN chmod +x wait-for.sh
COPY db/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
#ENTRYPOINT [ "/app/start.sh", "/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
