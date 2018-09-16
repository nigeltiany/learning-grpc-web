FROM golang:1.9-alpine

RUN apk add --no-cache git build-base curl
RUN curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o $GOPATH/bin/dep
RUN chmod +x $GOPATH/bin/dep

# Copy the local package files to the container's workspace.
WORKDIR /go/src/echo
ADD ./Gopkg.lock ./Gopkg.lock
ADD ./Gopkg.toml ./Gopkg.toml

RUN dep ensure -vendor-only
ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o echo

FROM scratch

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/src/echo /usr/bin/echo

ENTRYPOINT ["echo"]

EXPOSE 9080