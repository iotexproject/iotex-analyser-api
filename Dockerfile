FROM golang:1.23-bullseye AS builder

WORKDIR /app

# ENV GO111MODULE on
# ENV GOPROXY https://goproxy.cn

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apt-get update && apt-get install -y make gcc musl-dev git libc-dev build-essential

COPY . .
RUN go mod download
RUN make

FROM golang:1.23-bullseye

WORKDIR /app
COPY --from=builder /app/config.yml /app/
COPY --from=builder /app/iotex-analyser-api /app/
ENTRYPOINT ["./iotex-analyser-api"]