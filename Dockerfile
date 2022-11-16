


FROM golang:1.19-alpine as builder

RUN apk add --no-cache gcc g++

WORKDIR /build

ADD Hello.jpg ./
ADD config/local.toml ./
ADD config/token.txt ./
ADD DB/upgrade.db ./

ADD go.mod go.sum* ./

RUN go mod download

COPY . .

RUN go get -u gopkg.in/telebot.v3
RUN go get github.com/BurntSushi/toml@latest
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/sqlite
RUN go mod vendor

RUN go build -v -mod=vendor -o app main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /build/app /usr/local/bin/
COPY --from=builder /build/Hello.jpg ./

RUN mkdir config
COPY --from=builder /build/config/local.toml ./config/
COPY --from=builder /build/config/token.txt ./config/

RUN mkdir DB
COPY --from=builder /build/DB/upgrade.db ./DB/
RUN apk add sqlite

ARG PORT
ARG CONFIG

CMD /usr/local/bin/app -port=$PORT -config=$CONFIG