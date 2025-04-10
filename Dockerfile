FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v

RUN go build -o /app/cmd/site

FROM scratch

COPY --from=builder /app/cmd/site /site
COPY *.yml /
COPY templates/* /templates/

EXPOSE 8080
ENTRYPOINT ["/site"]