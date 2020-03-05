FROM golang:1.13

WORKDIR /go/src/patchwork
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# RUN mkdir -p /data/app
# ADD ./patchwork /data/app/
# RUN go build -o main .

EXPOSE 8080

CMD ["patchwork"]
