FROM golang:latest as builder
WORKDIR /go/src/peddecord-photo-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -o peddecord-photo-api

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/peddecord-photo-api/peddecord-photo-api .
CMD ["./peddecord-photo-api"]