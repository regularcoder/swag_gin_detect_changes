FROM golang:1.17

WORKDIR /go/src

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY detect_swagger_changes.sh detect_swagger_changes.sh

RUN chmod +x detect_swagger_changes.sh

ENTRYPOINT ["sh", "detect_swagger_changes.sh"]