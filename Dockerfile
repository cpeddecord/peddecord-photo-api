FROM golang:latest as builder
WORKDIR /go/src/peddecord-photo-api
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix -o peddecord-photo-api

FROM scratch
ARG COMMIT_REF
ARG BUILD_DATE
ENV COMMIT_REF=${COMMIT_REF} \
  BUILD_DATE=${BUILD_DATE}
WORKDIR /root/
COPY --from=builder /go/src/peddecord-photo-api/peddecord-photo-api .
CMD ["./peddecord-photo-api"]