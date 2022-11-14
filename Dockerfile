

FROM golang:1.19-alpine as builder

WORKDIR /build

ADD Hello.jpg ./

ADD go.mod go.sum* ./

RUN go mod download

COPY . .

RUN go get -u gopkg.in/telebot.v3
RUN go get github.com/BurntSushi/toml@latest
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/sqlite
RUN go mod vendor

# пока build сломался так как надо установить gcc для sqlite или настроить сразу mysql

RUN go build -v -mod=vendor -o app main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /build/app /usr/local/bin/
COPY --from=builder /build/Hello.jpg ./

ARG PORT

CMD /usr/local/bin/app -port=$PORT