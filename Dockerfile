FROM --platform=linux/amd64 golang:1.22

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct
ENV TZ=Asia/Shanghai

RUN apt-get update && apt-get install -y tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o goiot .

VOLUME ["/app/logs"]

ENTRYPOINT ["./iot"]
