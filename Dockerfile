FROM golang:1.14-alpine as builder

COPY . /car-pooling-challenge-sonercirit
WORKDIR /car-pooling-challenge-sonercirit
RUN go build

FROM alpine:3.8

# This Dockerfile is optimized for go binaries, change it as much as necessary
# for your language of choice.

RUN apk --no-cache add ca-certificates libc6-compat

EXPOSE 9091

COPY --from=builder /car-pooling-challenge-sonercirit/car-pooling-challenge-sonercirit /
COPY --from=builder /car-pooling-challenge-sonercirit/.env /

ENTRYPOINT [ "sh", "-c", "source /.env && /car-pooling-challenge-sonercirit" ]
