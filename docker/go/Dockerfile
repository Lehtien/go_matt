FROM golang:1.19-alpine

# ENV GO111MODULE=on

RUN apk update && \
    apk add emacs curl git
    # go get github.com/pilu/fresh

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]