FROM golang:latest as builder
WORKDIR /go/src/peddecord-photo-api
COPY . .
RUN (make api-build-ci)

FROM alpine:latest
ARG COMMIT_REF
ARG BUILD_DATE
ENV COMMIT_REF=${COMMIT_REF} \
  BUILD_DATE=${BUILD_DATE}
WORKDIR /root/
COPY --from=builder /go/src/peddecord-photo-api/peddecord-photo-api .
CMD ["./peddecord-photo-api"]