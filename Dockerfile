FROM golang:1.17
WORKDIR /app
RUN go get -u github.com/pressly/goose/v3/cmd/goose
RUN go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen
RUN apt update && apt install -y postgresql

RUN GO111MODULE=off go get github.com/oxequa/realize

CMD ["realize", "start"]
