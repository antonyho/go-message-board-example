FROM golang:1.13 AS build
WORKDIR /go/src/github.com/antonyho/go-message-board-example
COPY . .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN make

FROM scratch AS runtime
COPY --from=build /go/src/github.com/antonyho/go-message-board-example/messageboard-api-example /go/bin/
EXPOSE 8080/tcp
ENTRYPOINT ["/go/bin/messageboard-api-example"]
