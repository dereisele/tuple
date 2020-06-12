FROM golang:latest as build-env

WORKDIR /go/src/tuple
ADD . /go/src/tuple
RUN go get -d -v ./...
RUN go build -o /go/bin/tuple

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/tuple /
CMD ["/tuple"]
